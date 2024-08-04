package service

import (
	"common/grpcs/user/pb"
	"context"
	"errors"
	"fmt"
	"time"
	"users-microservice/gateway"
	"users-microservice/pkg/repo"
	"users-microservice/util"

	"github.com/samber/do"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthSvcs struct {
	pb.UnimplementedAuthServiceServer
	jwtManager     JwtManager
	userrepo       repo.UserRepo
	rolerepo       repo.RoleRepo
	capabilityrepo repo.CapabilityRepo
	emailgw        *gateway.EmailGw
}

func NewAuthSvcs(i *do.Injector) (*AuthSvcs, error) {
	return &AuthSvcs{
		jwtManager:     do.MustInvoke[JwtManager](i),
		userrepo:       do.MustInvoke[repo.UserRepo](i),
		rolerepo:       do.MustInvoke[repo.RoleRepo](i),
		capabilityrepo: do.MustInvoke[repo.CapabilityRepo](i),
		emailgw:        do.MustInvoke[*gateway.EmailGw](i),
	}, nil
}

// grpc method
func (a *AuthSvcs) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {

	email := req.GetEmail()
	password := req.GetPassword()

	user, err := a.userrepo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, status.Errorf(codes.Unknown, "something went wrong!")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(password)); err != nil {
		return nil, status.Errorf(codes.Unknown, "something went wrong!")
	}

	token, err := a.jwtManager.Generate(user, time.Duration(time.Hour*24), "users", "auth")
	if err != nil {
		return nil, err
	}

	return &pb.LoginRes{
		Token: token,
	}, nil

}

func (a *AuthSvcs) LoginWithProvider(ctx context.Context, req *pb.LoginWithProvReq) (*pb.LoginRes, error) {
	provider := req.GetProvider()
	token := req.GetToken()

	var user *pb.User
	var errc error

	if provider == "google" {
		payload, err := idtoken.Validate(context.Background(), token, "")
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}

		claims := payload.Claims

		email := claims["email"].(string)
		lastName := claims["family_name"].(string)
		firstName := claims["given_name"].(string)
		email_verified := claims["email_verified"].(bool)

		//check if the user exist
		user, errc = a.userrepo.GetUserByEmail(ctx, email)

		if errc != nil {
			if errc == util.ErrNotFound {
				//insert the user
				newUser := &pb.User{
					FirstName:         firstName,
					LastName:          lastName,
					Email:             email,
					EmailConfirmation: email_verified,
					Password:          "12333443fsfrf",
				}
				user, errc = a.userrepo.CreateUser(ctx, newUser)
				if errc != nil {
					return nil, status.Errorf(codes.Internal, err.Error())
				}

			} else {
				return nil, status.Errorf(codes.Internal, "error get user")
			}
		}

	} else if provider == "facebook" {

	}

	token, err := a.jwtManager.Generate(user, time.Duration(time.Hour*24), "users", "auth")
	if err != nil {
		return nil, err
	}

	return &pb.LoginRes{
		Token: token,
	}, nil
}

// any service can call this method to authorize user
func (a *AuthSvcs) Guard(ctx context.Context, req *pb.GuardReq) (*pb.GuardRes, error) {
	fmt.Println("invoke guard meth")
	restructions := req.GetRestrictions()
	//tobe removed later

	if len(restructions) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty restructions not allowed")
	}

	authStatus, err := a.IsUserCan(ctx, restructions...)

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "not authenticated")
	}

	return &pb.GuardRes{
		Authstatus: authStatus,
	}, nil
}

func (a *AuthSvcs) IsUserCan(ctx context.Context, caps ...string) (*pb.AuthStatus, error) {

	values, err := util.GetMdVal(ctx, "token")
	if err != nil {
		return nil, err
	}

	accessToken := values[0]

	claims, err := a.jwtManager.Verify(accessToken)
	if err != nil {
		return nil, errors.New("error verify token")
	}

	if len(caps) == 1 {
		if caps[0] == "auth" {
			return &pb.AuthStatus{
				IsAuthorized: true,
				UserId:       claims.UserId,
			}, nil
		} else {
			ok, err := a.UserHaveCapabilites(ctx, claims, caps[0])
			if err != nil {
				return nil, err
			} else {
				if ok {
					return &pb.AuthStatus{
						IsAuthorized: true,
						UserId:       claims.UserId,
					}, nil

				} else {
					return &pb.AuthStatus{
						IsAuthorized: false,
						UserId:       claims.UserId,
					}, nil
				}
			}

		}

	} else {
		capabilitiesOnly := []string{}
		for _, c := range caps {
			if c != "auth" {
				capabilitiesOnly = append(capabilitiesOnly, c)
			}
		}

		ok, err := a.UserHaveCapabilites(ctx, claims, capabilitiesOnly...)
		if err != nil {
			return nil, err
		} else {
			if ok {
				return &pb.AuthStatus{
					IsAuthorized: true,
					UserId:       claims.UserId,
				}, nil

			} else {
				return &pb.AuthStatus{
					IsAuthorized: false,
					UserId:       claims.UserId,
				}, nil
			}
		}

	}

}

func (a *AuthSvcs) UserHaveCapabilites(ctx context.Context, claims *UserClaims, capabilities ...string) (bool, error) {
	userId := claims.UserId
	userCapabilities, err := a.capabilityrepo.ReadUserCapabilities(ctx, userId)
	if err != nil {
		return false, err
	}

OUTER:
	for _, c := range capabilities {
		for _, related := range userCapabilities {
			if c == related {
				continue OUTER
			}
		}
		return false, nil
	}

	return true, nil
}

//

func (a *AuthSvcs) ForgotPassword(ctx context.Context, req *pb.ForgotPassReq) (*pb.ForgotPassRes, error) {
	email := req.GetEmail()

	//check if email exist
	user, err := a.userrepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error get user")
	}

	//generate forgot password token
	token, err := a.jwtManager.Generate(user, time.Duration(time.Hour), "users", "forgotpassword")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error generate token")
	}

	url := "something.com/resetpassword?token=" + token

	a.emailgw.SendEmail(ctx, url)

	return nil, nil
}

func (a *AuthSvcs) ResetPassword(ctx context.Context, req *pb.ResetPassReq) (*pb.ResetPassRes, error) {
	values, err := util.GetMdVal(ctx, "token")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}
	accessToken := values[0]

	claims, err := a.jwtManager.Verify(accessToken)

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	if claims.Access == "forgotpassword" {
		newPassword := req.GetPassword()
		if err := a.userrepo.ResetPassword(ctx, claims.UserId, newPassword); err != nil {
			return nil, status.Errorf(codes.Internal, "error reset password")
		}
	} else {
		return nil, status.Errorf(codes.Unknown, "error access token type")
	}

	return &pb.ResetPassRes{}, nil

}

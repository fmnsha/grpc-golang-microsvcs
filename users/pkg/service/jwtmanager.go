package service

import (
	"common/grpcs/user/pb"
	"fmt"
	"time"
	"users-microservice/pkg/util"

	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/do"
)

type UserClaims struct {
	jwt.RegisteredClaims
	UserId uint64
	Access string
}

type JwtManager interface {
	Generate(user *pb.User, expiresAt time.Duration, issuer, access string) (string, error)
	Verify(accesstoken string) (*UserClaims, error)
}

type jwtManager struct {
	secret string
}

func NewJwtManager(i *do.Injector) (JwtManager, error) {
	secret := util.GetEnv("secret", "")
	return &jwtManager{
		secret: secret,
	}, nil
}

func (j *jwtManager) Generate(user *pb.User, expiresAt time.Duration, issuer, access string) (string, error) {
	key := []byte(j.secret) /* Load key from somewhere, for example a file */
	claims := UserClaims{
		jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresAt)),
			Issuer:    issuer,
		},
		user.Id,
		access,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(key)
}

func (j *jwtManager) Verify(accesstoken string) (*UserClaims, error) {
	var claims UserClaims
	token, err := jwt.ParseWithClaims(accesstoken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*UserClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token claims")
	}

}

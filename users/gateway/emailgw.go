package gateway

import (
	"common/discovery"
	"context"
	"fmt"

	pb "common/grpcs/email/pb"

	"github.com/samber/do"
)

type EmailGw struct {
	registry *discovery.Registry
}

func NewEmailGw(i *do.Injector) (*EmailGw, error) {
	return &EmailGw{
		registry: do.MustInvoke[*discovery.Registry](i),
	}, nil
}

func (e *EmailGw) SendEmail(ctx context.Context, data string) (*pb.SendEmailRes, error) {

	fmt.Println("send email", data)

	conn, err := e.registry.GetConn(ctx, "emails")
	if err != nil {
		return nil, err
	}

	email := &pb.EmailMessage{
		To:      "feras@larsa.org",
		Subject: "test",
		Message: "test",
	}

	emailsvcs := pb.NewEmailServiceClient(conn)

	res, err := emailsvcs.SendEmail(ctx, &pb.SendEmailReq{
		Email: email,
	})

	if err != nil {
		return nil, err
	}

	return res, nil

}

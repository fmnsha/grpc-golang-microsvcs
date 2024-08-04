package models

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type Creds struct {
	Token  string
	Client string
}

func NewCreds(ctx context.Context) *Creds {

	c := &Creds{}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return c
	}

	if vals := md.Get("token"); len(vals) == 1 {
		c.Token = vals[0]
	}
	if vals := md.Get("client"); len(vals) == 1 {
		c.Client = vals[0]
	}

	return c
}

func (c *Creds) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"client": c.Client,
		"token":  c.Token,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security.
func (c Creds) RequireTransportSecurity() bool {
	return false
}

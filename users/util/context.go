package util

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/metadata"
)

type CtxKey int

const (
	_ CtxKey = iota
	Requser
)

func GetMdVal(ctx context.Context, key string) ([]string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}
	values := md.Get(key)
	if len(values) == 0 {
		return nil, fmt.Errorf("no values for key %s", key)
	}

	return values, nil
}

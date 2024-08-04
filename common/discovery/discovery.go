package discovery

import (
	"common/models"
	"context"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
)

type Registry struct {
	Client *clientv3.Client
}

func NewRegistry() (*Registry, error) {
	cli, cerr := clientv3.NewFromURL("http://localhost:2379")
	if cerr != nil {
		return nil, cerr
	}
	return &Registry{
		Client: cli,
	}, nil
}

func (r Registry) Register(ctx context.Context, serviceName, addr string) error {
	em, err := endpoints.NewManager(r.Client, serviceName)
	if err != nil {
		return err
	}
	if err := em.AddEndpoint(ctx, serviceName+"/", endpoints.Endpoint{Addr: addr}); err != nil {
		return err
	}
	return nil
}
func (r Registry) DeRegister(ctx context.Context, serviceName string) error {
	em, err := endpoints.NewManager(r.Client, serviceName)
	if err != nil {
		return err
	}
	if err := em.DeleteEndpoint(ctx, serviceName+"/"); err != nil {
		return err
	}
	return nil
}
func (r Registry) GetConn(ctx context.Context, serviceName string, dialOps ...grpc.DialOption) (*grpc.ClientConn, error) {

	fmt.Println(serviceName)

	cred := models.NewCreds(ctx)

	etcdResolver, err := resolver.NewBuilder(r.Client)
	if err != nil {
		return nil, err
	}

	fmt.Println(etcdResolver)
	opts := []grpc.DialOption{
		grpc.WithResolvers(etcdResolver),
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(cred),
	}
	opts = append(opts, dialOps...)
	conn, gerr := grpc.Dial("etcd:/"+serviceName, opts...)
	if gerr != nil {
		return nil, gerr
	}

	return conn, nil
}

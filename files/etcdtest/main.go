package main

import (
	"context"
	"fmt"
	"log"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

func main() {
	cli, cerr := clientv3.NewFromURL("http://localhost:2379")
	if cerr != nil {
		log.Fatal(cerr)
	}
	em, _ := endpoints.NewManager(cli, "foo/bar/my-service")
	em.AddEndpoint(context.TODO(), "foo/bar/my-service/e1", endpoints.Endpoint{Addr: "1.2.3.4"})

	res, _ := cli.Get(context.Background(), "foo/bar/my-service/e1")

	fmt.Println(res)

	// etcdResolver, _ := resolver.NewBuilder(cli)

	// fmt.Printf("%s", etcdResolver)

	// // etcdResolver, err := resolver.NewBuilder(cli)
	// conn, gerr := grpc.Dial("etcd:///foo/bar/my-service", grpc.WithResolvers(etcdResolver))

	// if gerr != nil{

	// }
}

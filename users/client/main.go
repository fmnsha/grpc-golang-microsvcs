package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"common/grpcs/user/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:3000", "the address to connect to")

type MyData struct {
	client string
}

func (m *MyData) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"client": m.client}, nil
}
func (m *MyData) RequireTransportSecurity() bool {
	return false
}

func main() {
	flag.Parse()

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("client1", 1)
	go CallClient1("microsvcs", 2)
	go CallClient1("microsvcs", 2)

	time.Sleep(time.Second * 4)
}

func CallClient1(client string, id uint64) {
	opts := []grpc.DialOption{
		// In addition to the following grpc.DialOption, callers may also use
		// the grpc.CallOption grpc.PerRPCCredentials with the RPC invocation
		// itself.
		// See: https://godoc.org/google.golang.org/grpc#PerRPCCredentials
		grpc.WithPerRPCCredentials(&MyData{
			client: client,
		}),
		// oauth.TokenSource requires the configuration of transport
		// credentials.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	rgc := pb.NewUserServiceClient(conn)

	res, err := rgc.ReadUser(context.Background(), &pb.ReadUserReq{
		Id: id,
	})

	if err != nil {
		log.Fatal(err)
	}

	t := res.User.CreatedAt.AsTime()

	fmt.Println(res)
	fmt.Println("the time is", t)
}

// fetchToken simulates a token lookup and omits the details of proper token
// acquisition. For examples of how to acquire an OAuth2 token, see:
// https://godoc.org/golang.org/x/oauth2
// func fetchToken() *oauth2.Token {
// 	return &oauth2.Token{
// 		AccessToken: "some-secret-token",
// 	}
// }

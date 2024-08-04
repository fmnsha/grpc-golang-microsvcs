package main

import (
	"common/grpcs/file/pb"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
)

type Mt struct{}

func (m Mt) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {

	return map[string]string{
		"client": "microsvcs",
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security.
func (m Mt) RequireTransportSecurity() bool {
	return false
}

func main() {

	opts := []grpc.DialOption{
		// In addition to the following grpc.DialOption, callers may also use
		// the grpc.CallOption grpc.PerRPCCredentials with the RPC invocation
		// itself.
		// See: https://godoc.org/google.golang.org/grpc#PerRPCCredentials
		grpc.WithPerRPCCredentials(Mt{}),
		grpc.WithInsecure(),
		// oauth.TokenSource requires the configuration of transport
		// credentials.
		//grpc.WithTransportCredentials(creds),
	}
	conn, err := grpc.Dial("localhost:3001", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewFileServiceClient(conn)

	stream, err := c.Upload(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("./img.webp")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	buf := make([]byte, 1000)

	for {
		num, err := file.Read(buf)

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		chunk := buf[:num]

		if err := stream.Send(&pb.FileUploadReq{FileName: file.Name(), Chunk: chunk}); err != nil {
			log.Fatal(err)
		}

	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println(res)

}

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("server.crt", "server.grpc.io")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := NewHelloServiceClient(conn)

	reply, err := client.Hello(context.Background(), &String{Value: "hello it is me"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetValue())

	stream, err := client.Channel(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for {
			if err := stream.Send(&String{Value: " hi"}); err != nil {
				fmt.Println(err)
			}

			time.Sleep(time.Second * 5)
		}

	}()

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println(err)
		}

		fmt.Println(reply.GetValue())

	}

}

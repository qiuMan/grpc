package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := NewPubsubServiceClient(conn)

	_, err = client.Publish(context.Background(), &String{Value: "golang: hello Go"})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Publish(context.Background(), &String{Value: "docker: hello docker"})
	if err != nil {
		log.Fatal(err)
	}

}

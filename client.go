package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"

	pbHello "pmapp/api/hello"
)

var host string = "server"

const port uint32 = 9001

func install() error {
	var err error

	if len(os.Args) > 1 {
		host = os.Args[1]
	}
	address := fmt.Sprintf("%s:%d", host, port)
	log.Println("start client for", address)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("error dial:", err)
		return err
	}
	defer conn.Close()

	req := pbHello.InstallRequest{
		Hostname: "localhost",
		Port:     12345,
	}

	pbClient := pbHello.NewHelloClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := pbClient.Install(ctx, &req)
	if err != nil {
		return err
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			return err
		}
		if res.Done {
			log.Println("installation is finished!")
			break
		}
		log.Println("installation is incomplete")
	}
	return err
}

func main() {
	log.Println("start client")
	err := install()
	if err != nil {
		fmt.Println("client error:", err)
	}
}

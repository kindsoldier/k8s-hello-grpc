package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pbHello "pmapp/api/hello"

	"google.golang.org/grpc"
)

const (
	port = 9001
)

type HelloServer struct {
	pbHello.UnimplementedHelloServer
}

func NewHelloServer() *HelloServer {
	return &HelloServer{}
}

func (helloServ *HelloServer) Install(req *pbHello.InstallRequest, stream pbHello.Hello_InstallServer) error {
	var err error
	log.Println("install request:", req)

	var wg sync.WaitGroup
	doneChan := make(chan bool)

	aliveFunc := func() {
		defer wg.Done()
		for {
			select {
			case <-doneChan:
				log.Println("cancel alive func")
				return
			default:
				time.Sleep(1 * time.Second)
			}
			log.Println("send alive response")
			intermRes := pbHello.InstallResult{
				Done: false,
			}
			err := stream.Send(&intermRes)
			if err != nil {
				log.Println("send alive error")
				return
			}
		}
	}
	wg.Add(1)
	go aliveFunc()

	time.Sleep(120 * time.Second)
	log.Println("installation is finished!")

	doneChan <- true
	wg.Wait()

	doneRes := pbHello.InstallResult{
		Done: true,
	}
	err = stream.Send(&doneRes)
	if err != nil {
		return err
	}

	return err
}

func main() {
	var err error

	addresses, err := net.InterfaceAddrs()

	if err == nil {
		for _, addr := range addresses {
			fmt.Println("service ip addess:", addr)

		}
	}
	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServ := grpc.NewServer()
	pbHello.RegisterHelloServer(grpcServ, NewHelloServer())

	log.Printf("server listening at %d", port)
	err = grpcServ.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

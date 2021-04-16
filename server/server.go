package main

import (
	"fmt"
	"grpc_mongo/proto/blog"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	// if we crash the go code, we can now the line
	// log trick
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// listener
	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("%v", err)
	}

	// server
	s := grpc.NewServer()
	// register blog service
	blog.RegisterBlogServiceServer(s, &server{})

	// mongo_client := mongo_connection(5 * time.Second)
	// log.Println(mongo_client)

	// run server with go routine
	go func() {
		fmt.Println("Server starting ...")
		startErr := s.Serve(lis)
		if startErr != nil {
			log.Fatalf("Error when starting the server ...")
		}
	}()

	// wait control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// block unitl signal is received
	<-ch

	// gracefully shutdown
	fmt.Println("Stopping the server ...")
	s.Stop()
	fmt.Println("Stopping the listener")
	lis.Close()
	fmt.Println("Server shutdown gracefully")
}

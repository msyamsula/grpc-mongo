package main

import (
	"fmt"
	"grpc_mongo/connection"
	"grpc_mongo/proto/blog"
	"grpc_mongo/service"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// if we crash the go code, we can now the line
	// log trick
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	uri := "mongodb://mongo:mongo@localhost:27017"
	dbName := "grpc"
	connection.MongoConnection(uri, dbName, 5*time.Second)

	// listener
	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("%v", err)
	}

	// server
	s := grpc.NewServer()
	// register blog service
	blog.RegisterBlogServiceServer(s, &service.Server{})

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
	fmt.Println("stopping mongo client")
	fmt.Println("Stopping the server ...")
	s.Stop()
	fmt.Println("Stopping the listener")
	lis.Close()
	fmt.Println("Server shutdown gracefully")
}

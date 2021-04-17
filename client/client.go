package main

import (
	"context"
	"fmt"
	"grpc_mongo/proto/blog"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("something wrong")
		log.Fatal(err)
	}

	c := blog.NewBlogServiceClient(conn)

	req := &blog.CreateBlogRequest{
		AuthorId: "4",
		Title:    "Syamsul",
		Content:  "Cepak Cepek",
	}

	resp, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	println(resp.GetStatus())

	getAllReq := &blog.GetAllBlogRequest{}
	stream, serr := c.GetAllBlog(context.TODO(), getAllReq)
	if serr != nil {
		log.Fatal(serr)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println(stream.CloseSend())
			break
		} else if err != nil {
			log.Fatal(err)
		}

		author_id := msg.GetAuthorId()
		content := msg.GetContent()
		title := msg.GetContent()

		fmt.Printf("get data --> author_id: %v, title: %v, content: %v\n", author_id, title, content)
	}
}

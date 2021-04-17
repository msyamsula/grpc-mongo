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

	// create
	req := &blog.CreateBlogRequest{
		AuthorId: "7",
		Title:    "Muhammad",
		Content:  "Competitive Programming",
	}

	resp, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	println(resp.GetStatus())

	// update call
	updateReq := &blog.UpdateBlogRequest{
		AuthorId:    "",
		Tile:        "Arifin",
		Content:     "",
		NewAuthorId: "2",
		NewTile:     "Arifin",
		NewContent:  "Tidur",
	}

	ures, uerr := c.UpdateBlog(context.Background(), updateReq)

	if uerr != nil {
		log.Println(uerr)
		log.Println(ures.GetResponse())
		return
	}

	log.Println(ures.GetResponse())

	// get all
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
		title := msg.GetTitle()

		fmt.Printf("get data --> author_id: %v, title: %v, content: %v\n", author_id, title, content)
	}

	// clean up
	cleanReq := &blog.CleanBlogRequest{}
	cresp, cerr := c.CleanBlog(context.Background(), cleanReq)
	if cerr != nil {
		log.Println(err)
		log.Println(cresp.GetResponse())
		return
	}

	log.Println(cresp.GetResponse())
}

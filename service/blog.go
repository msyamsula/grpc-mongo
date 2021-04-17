package service

import (
	"context"
	"fmt"
	"grpc_mongo/connection"
	"grpc_mongo/model"
	"grpc_mongo/proto/blog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (Server) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.CreateBlogResponse, error) {
	fmt.Print("create blog is called\n")
	author_id := req.GetAuthorId()
	title := req.GetTitle()
	content := req.GetContent()

	data := model.Blog{
		ID:         primitive.NewObjectID(),
		Author_id:  author_id,
		Title:      title,
		Content:    content,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	blogCollection := connection.MongoDB.Collection("Blog")

	_, err := blogCollection.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}

	res := &blog.CreateBlogResponse{
		Status: "success",
	}

	return res, nil
}

func (*Server) GetAllBlog(req *blog.GetAllBlogRequest, stream blog.BlogService_GetAllBlogServer) error {

	fmt.Println("get all is called")
	blogCollection := connection.MongoDB.Collection("Blog")
	ctx := context.TODO()
	cursor, err := blogCollection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var data model.Blog
		err := cursor.Decode(&data)
		if err != nil {
			return err
		}

		res := &blog.GetAllBlogResponse{
			AuthorId: data.Author_id,
			Title:    data.Title,
			Content:  data.Content,
		}

		stream.Send(res)
	}

	return nil
}

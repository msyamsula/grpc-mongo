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

func (*Server) CleanBlog(ctx context.Context, req *blog.CleanBlogRequest) (*blog.CleanBlogResponse, error) {
	fmt.Println("clean up database, can't be undone, you should've consider this act first")
	blogCollection := connection.MongoDB.Collection("Blog")
	result, err := blogCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		resp := &blog.CleanBlogResponse{
			Response: "something wrong when deleting",
		}
		return resp, err
	}

	resp := &blog.CleanBlogResponse{
		Response: "succesfully clean up database",
	}

	fmt.Printf("delete: %v row(s)\n", result.DeletedCount)

	return resp, nil
}

func (*Server) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.UpdateBlogResponse, error) {
	author_id := req.AuthorId
	title := req.Tile
	content := req.Content

	updateFilter := bson.M{}

	if author_id != "" {
		updateFilter["author_id"] = author_id
	}

	if title != "" {
		updateFilter["title"] = title
	}

	if content != "" {
		updateFilter["content"] = content
	}

	if author_id == "" && title == "" && content == "" {
		resp := &blog.UpdateBlogResponse{
			Response: "no filter given",
		}

		return resp, nil
	}

	newAuthorId := req.NewAuthorId
	newTitle := req.NewTile
	newContent := req.NewContent

	updateValue := bson.M{}
	if author_id != "" {
		updateValue["author_id"] = newAuthorId
	} else {
		updateValue["author_id"] = author_id
	}

	if title != "" {
		updateValue["title"] = newTitle
	} else {
		updateValue["title"] = title
	}

	if content != "" {
		updateValue["content"] = newContent
	} else {
		updateValue["content"] = content
	}

	blogCollection := connection.MongoDB.Collection("Blog")

	result, err := blogCollection.ReplaceOne(ctx, updateFilter, updateValue)
	if err != nil {
		resp := &blog.UpdateBlogResponse{
			Response: "somethings wrong when updating",
		}
		return resp, err
	}

	resp := &blog.UpdateBlogResponse{
		Response: "succesfully update database",
	}

	fmt.Printf("%v row(s) was updated", result.ModifiedCount)
	return resp, nil
}

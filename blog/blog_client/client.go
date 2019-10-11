package main

import (
	"../blogpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Blog client")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatalf("Could not connect %v\n", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// create blog
	blogId := createBlog(c)
	readBlog(c, blogId)
	updateBlog(c, blogId)
	deleteBlog(c, blogId)

	createBlog(c)
	createBlog(c)
	listBlog(c)
}

func createBlog(c blogpb.BlogServiceClient) string {
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Muhammed",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}
	fmt.Printf("Blog has been created: %v\n", createBlogRes)
	return createBlogRes.GetBlog().GetId()
}

func readBlog(c blogpb.BlogServiceClient, blogId string) {
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: "5bdc29e661b75adcac496cf4",
	})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v\n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogId}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error happened while reading: %v\n", readBlogErr)
	}
	fmt.Printf("Blog was read: %v\n", readBlogRes)
}

func updateBlog(c blogpb.BlogServiceClient, blogId string) {
	newBlog := &blogpb.Blog{
		Id:       blogId,
		AuthorId: "Changed Author",
		Title:    "My First Blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})

	if updateErr != nil {
		fmt.Printf("Error happened while updating: %v\n", updateErr)
	}
	fmt.Printf("Blog was updated: %v\n", updateRes)
}

func deleteBlog(c blogpb.BlogServiceClient, blogId string) {
	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{
		BlogId: blogId,
	})

	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v\n", deleteErr)
	}
	fmt.Printf("Blog was deleted: %v\n", deleteRes)
}

func listBlog(c blogpb.BlogServiceClient) {
	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Printf("Error while calling ListBlog RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetBlog())
	}
}

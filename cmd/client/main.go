package main

import (
	"context"
	"flag"
	"log"

	"github.com/khdoba2000/grpc-articles/pkg/proto"

	"google.golang.org/grpc"
)

func main() {
	//serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("Dial server : %s", "2000")

	conn, err := grpc.Dial("localhost:2000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Cannot dial server: ", err)
	}

	ArticleClient := proto.NewArticleServiceClient(conn)

	article := &proto.Article{
		Title:   "title55 sdasdadjk",
		Desc:    "description55 (desssss",
		Content: "content55 conteeeeeeeeent",
	}

	req := proto.CreateArticleRequest{
		Article: article,
	}

	res, err := ArticleClient.Create(context.Background(), &req)
	if err != nil {
		log.Fatal("Cannot Create article: ", err)
		return
	}

	log.Printf("created article with id: %d", res.Id)
}

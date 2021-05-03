package main

import (
	"fmt"
	"log"
	"net"

	"github.com/khdoba2000/grpc-articles/config"
	"github.com/khdoba2000/grpc-articles/pkg/proto"
	"github.com/khdoba2000/grpc-articles/service"

	"google.golang.org/grpc"
)

func main() {
	// port := flag.Int("port", 0, "the server port")
	// flag.Parse()
	// log.Printf("Start server on port: %d", *port)

	db := config.Connect()
	ArticleServiceServer := service.NewArticleServiceServer(db)

	grpcServer := grpc.NewServer()
	proto.RegisterArticleServiceServer(grpcServer, ArticleServiceServer)

	//address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", "0.0.0.0:2000")
	fmt.Print("Started server on 2000")
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}

}

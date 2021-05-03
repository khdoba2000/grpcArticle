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
    // //Create
	// article := &proto.Article{
	// 	Title:   "title999 sdasdadjk",
	// 	Desc:    "descripti9995 (desssss",
	// 	Content: "content999 conteeeeeeeeent",
	// }

	// req := proto.CreateArticleRequest{
	// 	Article: article,
	// }

	// res, err := ArticleClient.Create(context.Background(), &req)
	// if err != nil {
	// 	log.Fatal("Cannot Create article: ", err)
	// 	return
	// }

	// log.Printf("created article with id: %d", res.Id)


	// // Read
	// req2 := proto.ReadArticleRequest{
	// 	Id: 2,
	// }
	// res2, err := ArticleClient.Read(context.Background(), &req2)
	// if err != nil {
	// 	log.Fatalf("Read failed: %v", err)
	// }
	// log.Printf("Read result: ReadArticle wit ID<%d>, Title:<%s>\n\n", res2.Article.Id, res2.Article.Title)


	// // Update
	// req3 := proto.UpdateArticleRequest{
	// 	Article: &proto.Article{
	// 		Id:      res2.Article.Id,
	// 		Title:   res2.Article.Title,
	// 		Desc:    res2.Article.Desc + " + updated",
	// 		Content: res2.Article.Content,
	// 	},
	// }
	// res3, err := ArticleClient.Update(context.Background(), &req3)
	// if err != nil {
	// 	log.Fatalf("Update failed: %v", err)
	// }
	// log.Printf("Update result: <%+v>\n\n", res3)



	// Call ReadAll
	req4 := proto.ReadAllArticleRequest{}
	res4, err := ArticleClient.ReadAll(context.Background(), &req4)
	if err != nil {
		log.Fatalf("	ReadAll failed: %v", err)
	}
	log.Printf("ReadAll result: <%+v>\n\n", res4)

	
	// Delete
	// req5 := proto.DeleteArticleRequest{
	// 	Id:  6,
	// }
	// res5, err := ArticleClient.Delete(context.Background(), &req5)
	// if err != nil {
	// 	log.Fatalf("Delete failed: %v", err)
	// }
	// log.Printf("Delete result: <%+v>\n\n", res5)

}

package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/khdoba2000/grpc-articles/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ArticleServer struct {
	db *sql.DB
	proto.UnimplementedArticleServiceServer
}

func NewArticleServiceServer(db *sql.DB) *ArticleServer {
	return &ArticleServer{
		db: db,
	}
}

// connect returns SQL database connection from the pool
func (server *ArticleServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := server.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

//Create Article
func (server *ArticleServer) Create(
	ctx context.Context,
	req *proto.CreateArticleRequest,
) (*proto.CreateArticleResponse, error) {
	article := req.Article
	log.Printf("Received a CREATE REQUEST")

	if article.Id < 0 {
		// check if it is a valid id
		return nil, errors.New("Not valid id for Article")
	}

	// get SQL connection from pool
	database, err := server.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	// insert Article entity data
	// dynamic
	// 	create sequence article_id_seq;
	// alter table ararticle2 alter id set default nextval('article_id_seq');
	// Select setval('player_id_seq', 2000051 );
	insertDynStmt := `insert into "article2" ("title", "descr", "content") values($1, $2, $3) returning "id"`

	LastInsertId := 0
	err = database.QueryRowContext(ctx, insertDynStmt, article.Title, article.Desc, article.Content).Scan(&LastInsertId)

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into articles-> "+err.Error())
	}


	return &proto.CreateArticleResponse{
		Id: int32(LastInsertId),
	}, nil
}

//Read Article
func (server *ArticleServer) Read(
	ctx context.Context,
	req *proto.ReadArticleRequest,
) (*proto.ReadArticleResponse, error) {

	if req.Id < 0 {
		// check if it is a valid id
		return nil, errors.New("Invalid ID error, the givem ID is negative")
	}

	// get SQL connection from pool
	database, err := server.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	// Query entity data
	// dynamic
	insertDynStmt := `SELECT ("id", "title", "descr", "content") FROM "article2" WHERE id=$1`

	rows, err := database.QueryContext(ctx, insertDynStmt, req.Id)

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to GET from articles-> "+err.Error())
	}

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Articles-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Article with ID='%d' is not found",
			req.Id))
	}

	// get Article data
	var article proto.Article
	if err := rows.Scan(&article.Id, &article.Title, &article.Desc, &article.Content); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ToDo row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ToDo rows with ID='%d'",
			req.Id))
	}

	return &proto.ReadArticleResponse{
		Article: &article,
	}, nil
}

// Update Article
func (server *ArticleServer) Update(ctx context.Context, req *proto.UpdateArticleRequest) (*proto.UpdateArticleResponse, error) {

	// get SQL connection from pool
	c, err := server.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// update ToDo

	res, err := c.ExecContext(ctx, `update "article2" set "title"=$2, "descr"=$3, "content"=$4 where "id"=$1`,
		req.Article.Title, req.Article.Desc, req.Article.Content, req.Article.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update ToDo-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Article with ID='%d' is not found",
			req.Article.Id))
	}
	return &proto.UpdateArticleResponse{
		Updated: int32(rows),
	}, nil
}

// Delete todo task
func (server *ArticleServer) Delete(ctx context.Context, req *proto.DeleteArticleRequest) (*proto.DeleteArticleResponse, error) {

	// get SQL connection from pool
	c, err := server.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// delete Article
	res, err := c.ExecContext(ctx, `delete from "article2" where id=$1`, req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete ToDo-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Article with ID='%d' is not found",
			req.Id))
	}

	return &proto.DeleteArticleResponse{
		Deleted: int32(rows),
	}, nil
}

// Read all todo tasks
func (server *ArticleServer) ReadAll(ctx context.Context, req *proto.ReadAllArticleRequest) (*proto.ReadAllArticleResponse, error) {

	// get SQL connection from pool
	c, err := server.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	// get ToDo list
	rows, err := c.QueryContext(ctx, `SELECT "id", "title", "descr", "content" FROM "article2"`)

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ToDo-> "+err.Error())
	}
	defer rows.Close()

	list := []*proto.Article{}
	for rows.Next() {
		article := new(proto.Article)
		if err := rows.Scan(&article.Id, &article.Title, &article.Desc); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Article row-> "+err.Error())
		}

		list = append(list, article)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Article-> "+err.Error())
	}

	return &proto.ReadAllArticleResponse{
		Articles: list,
	}, nil
}

func mustEmbedUnimplementedArticleServiceServer() {

}

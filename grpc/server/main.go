package main

import (
	"context"
	"log"
	"net"

	"github.com/kkoehler/golang/grpc/comment"
	grpc "google.golang.org/grpc"
)

type CommentsServerImpl struct {
}

func (cs CommentsServerImpl) CreateComment(ctx context.Context, newComment *comment.Comment) (*comment.Void, error) {

	log.Printf("server called with comment: %s", newComment.Message)
	return &comment.Void{}, nil

}

func (cs CommentsServerImpl) GetComments(context.Context, *comment.Void) (*comment.CommentList, error) {

	log.Println("returning list")
	list := &comment.CommentList{}

	list.Comments = append(list.Comments, &comment.Comment{Message: "eins"})
	list.Comments = append(list.Comments, &comment.Comment{Message: "zwei"})

	return list, nil
}

func main() {

	commentsImpl := CommentsServerImpl{}

	srv := grpc.NewServer()
	comment.RegisterCommentsServer(srv, commentsImpl)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("could not open tcp port %v", err)
	}

	log.Println("starting server")

	err = srv.Serve(listener)
	if err != nil {
		log.Fatalf("stopped server %v", err)
	}

}

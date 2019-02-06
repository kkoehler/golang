package main

import (
	"context"
	"log"

	"github.com/kkoehler/golang/grpc/comment"
	"google.golang.org/grpc"
)

func main() {

	ctx := context.Background()
	con, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not open connection %v", err)
	}

	client := comment.NewCommentsClient(con)

	commentToCreate := comment.Comment{Message: "Hello World", Who: "kristian"}

	_, err = client.CreateComment(ctx, &commentToCreate)
	if err != nil {
		log.Fatalf("could not call %v", err)
	}

	list, err := client.GetComments(ctx, &comment.Void{})
	if err != nil {
		log.Fatalf("could not read list %v", err)
	}

	for _, entry := range list.Comments {
		log.Println(entry)
	}

}

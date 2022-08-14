package main

import (
	"context"
	"fmt"

	"github.com/thatmobiledude/go-practice/internal/comment"
	"github.com/thatmobiledude/go-practice/internal/db"
)

// Run - is going to be responsible for the instantiation and startup of our go application
func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}
	fmt.Println("successfully connected and pinged database")

	cmtService := comment.NewService(db)

	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "71c5d074-b6cf-11ec-b909-0242ac120002",
			Slug:   "manual-test",
			Author: "Bender Rodriguez",
			Body:   "Eat my shiny metal a**",
		},
	)

	fmt.Println(cmtService.GetComment(
		context.Background(),
		"71c5d074-b6cf-11ec-b909-0242ac120002",
	))

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"

	"github.com/thatmobiledude/go-practice/internal/comment"
	transportHttp "github.com/thatmobiledude/go-practice/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

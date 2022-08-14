package comment

import (
	"context"
	"fmt"
)

// Comment - a respresentation of the comment structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - an interface to define all methods that our service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - is the struct in which all our logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new service
// using Constructors and composite literals over 'new' keyword
// this allows flexibility to initialize fields
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetComment - returns the comment with the associated userId or empty comment with error
// uses the service pointer to call the interfaces' GetComment function to retrieve the comment or error.
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}
	return cmt, nil
}

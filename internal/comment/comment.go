package comment

import (
	"context"
	"errors"
	"fmt"
)

// error definitions to not expose any db implementation details to our clients
var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrPostingComment  = errors.New("failed to post comment by id")
	ErrUpdatingComment = errors.New("failed to update comment by id")
	ErrDeletingComment = errors.New("failed to delete comment by id")
	ErrNotImplemented  = errors.New("not implemented")
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
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
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
		fmt.Println(err) // would usually log this to log app (e.g. DataDog, GrayLog, etc.)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(
	ctx context.Context,
	ID string,
	updatedCmt Comment,
) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, ID, updatedCmt)
	if err != nil {
		fmt.Println("error updating comment")
		return Comment{}, ErrUpdatingComment
	}
	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	err := s.Store.DeleteComment(ctx, id)
	if err != nil {
		fmt.Println("error deleting comment")
		return ErrDeletingComment
	}
	return nil
}

func (s *Service) PostComment(ctx context.Context, comment Comment) (Comment, error) {
	cmt, err := s.Store.PostComment(ctx, comment)
	if err != nil {
		return Comment{}, ErrPostingComment
	}
	return cmt, nil
}

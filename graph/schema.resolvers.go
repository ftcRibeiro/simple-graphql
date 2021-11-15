package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/ftcRibeiro/simple-graphql/graph/generated"
	"github.com/ftcRibeiro/simple-graphql/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Name:        input.Name,
		Description: input.Description,
		ID:          fmt.Sprintf("T%d", rand.Int()),
		UserID:      input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		Name:  input.Name,
		Email: input.Email,
		ID:    fmt.Sprintf("T%d", rand.Int()),
	}
	return r.UserRepo.CreateUser(user)
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	response, err := r.TodoRepo.GetTodos()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	response, err := r.UserRepo.GetUsers()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *todoResolver) Name(ctx context.Context, obj *model.Todo) (string, error) {
	name := obj.Name
	return name, nil
}
func (r *todoResolver) Description(ctx context.Context, obj *model.Todo) (string, error) {
	description := obj.Description
	return description, nil
}

package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/ftcRibeiro/simple-graphql/graph/model"
	"github.com/ftcRibeiro/simple-graphql/postgres"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos    []*model.Todo
	users    []*model.User
	UserRepo postgres.UserRepo
	TodoRepo postgres.TodoRepo
}

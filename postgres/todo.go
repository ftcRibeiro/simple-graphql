package postgres

import (
	"github.com/ftcRibeiro/simple-graphql/graph/model"
	"github.com/go-pg/pg/v10"
)

type TodoRepo struct {
	DB *pg.DB
}

func (t *TodoRepo) GetTodos() ([]*model.Todo, error) {
	var todos []*model.Todo
	err := t.DB.Model(&todos).Select()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *TodoRepo) CreateTodo(todo *model.Todo) (*model.Todo, error) {
	_, err := t.DB.Model(todo).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return todo, err
}

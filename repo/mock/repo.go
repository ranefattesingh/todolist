package mock

import (
	"context"

	"github.com/ranefattesingh/todolist/core"
)

var todos core.TodoItems = core.TodoItems{
	&core.TodoItem{
		ID:          1,
		Title:       "title 1",
		Description: "description 1",
	},
	&core.TodoItem{
		ID:          2,
		Title:       "title 2",
		Description: "description 2",
	},
	&core.TodoItem{
		ID:          3,
		Title:       "title 3",
		Description: "description 3",
	},
}

type myRepo struct {
}

func (m *myRepo) GetAll(ctx context.Context) (*core.TodoItems, error) {
	return &todos, nil
}

func (m *myRepo) AddTodo(ctx context.Context, item *core.TodoItem) error {
	return nil
}

func (m *myRepo) UpdateTodo(ctx context.Context, id int, item *core.TodoItem) error {
	return nil
}

func (m *myRepo) DeleteTodo(ctx context.Context, id int) error {
	return nil
}

func NewRepo() (*myRepo, error) {
	return &myRepo{}, nil
}

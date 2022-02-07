package core

import "context"

type TodoRepo interface {
	GetAll(ctx context.Context) (*TodoItems, error)
	AddTodo(ctx context.Context, item *TodoItem) error
	UpdateTodo(ctx context.Context, id int, item *TodoItem) error
	DeleteTodo(ctx context.Context, id int) error
}

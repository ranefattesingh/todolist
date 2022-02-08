package core

import "context"

type TodoService interface {
	GetAll(ctx context.Context) (*TodoItems, error)
	UpdateTodo(ctx context.Context, id int, item *TodoItem) error
	AddTodo(ctx context.Context, item *TodoItem) error
	UpdateStatus(ctx context.Context, id int) error
	DeleteTodo(ctx context.Context, id int) error
}

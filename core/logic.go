package core

import (
	"context"
)

type todoService struct {
	r TodoRepo
}

func (t *todoService) GetAll(ctx context.Context) (*TodoItems, error) {
	return t.r.GetAll(ctx)
}

func (t *todoService) UpdateTodo(ctx context.Context, id int, item *TodoItem) error {
	return t.r.UpdateTodo(ctx, id, item)
}

func (t *todoService) AddTodo(ctx context.Context, item *TodoItem) error {
	return t.r.AddTodo(ctx, item)
}

func (t *todoService) UpdateStatus(ctx context.Context, id int) error {
	return t.r.UpdateStatus(ctx, id)
}

func (t *todoService) DeleteTodo(ctx context.Context, id int) error {
	return t.r.DeleteTodo(ctx, id)
}

func NewTodoService(r TodoRepo) *todoService {
	return &todoService{r}
}

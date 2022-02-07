package psql

import (
	"context"

	"github.com/ranefattesingh/todolist/core"
)

func (m *myRepo) GetAll(ctx context.Context) (*core.TodoItems, error) {
	row := m.db.QueryRow("SELECT COUNT(Id) FROM todo_item")

	var count int
	if err := row.Scan(&count); err != nil {
		return nil, row.Err()
	}

	todos := make(core.TodoItems, count)
	rows, err := m.db.Query("SELECT Id, Title, Description FROM todo_item")
	if err != nil {
		return nil, err
	}

	index := 0
	for rows.Next() {
		todo := &core.TodoItem{}
		rows.Scan(&todo.ID, &todo.Title, &todo.Description)
		todos[index] = todo
	}

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

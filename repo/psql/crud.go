package psql

import (
	"context"
	"fmt"
	"time"

	"github.com/ranefattesingh/todolist/core"
)

func (m *myRepo) GetAll(ctx context.Context) (*core.TodoItems, error) {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	row := m.db.QueryRowContext(queryCtx, "SELECT COUNT(Id) FROM todo_item")

	var count int
	if err := row.Scan(&count); err != nil {
		return nil, row.Err()
	}

	todos := make(core.TodoItems, count)
	rows, err := m.db.QueryContext(queryCtx, "SELECT Id, Title, Description, Status FROM todo_item")
	if err != nil {
		return nil, err
	}

	index := 0
	for rows.Next() {
		todo := &core.TodoItem{}
		rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status)
		todos[index] = todo
		index++
	}

	return &todos, nil
}

func (m *myRepo) UpdateTodo(ctx context.Context, id int, item *core.TodoItem) error {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	q := fmt.Sprintf(`
		UPDATE todo_item
		SET Title = '%s',
		Description = '%s'
		Status = %v
		WHERE
		Id = %d
	`, item.Title, item.Description, item.Status, id)
	_, err := m.db.ExecContext(queryCtx, q)

	if err != nil {
		return err
	}

	return nil
}

func (m *myRepo) AddTodo(ctx context.Context, item *core.TodoItem) error {
	queryCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	q := fmt.Sprintf("INSERT INTO todo_item(Id, Title, Description) VALUES(%d, '%s', '%s')",
		item.ID,
		item.Title,
		item.Description,
	)

	_, err := m.db.ExecContext(queryCtx, q)
	if err != nil {
		return err
	}

	return nil
}

func (m *myRepo) UpdateStatus(ctx context.Context, id int) error {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	row := m.db.QueryRow(fmt.Sprintf("SELECT COALESCE(Status, FALSE) FROM todo_item WHERE Id=%d", id))

	var status bool
	if err := row.Scan(&status); err != nil {
		return err
	}

	q := fmt.Sprintf("UPDATE todo_item SET Status = %t WHERE Id = %d", !status, id)
	_, err := m.db.ExecContext(queryCtx, q)
	if err != nil {
		return err
	}

	return nil
}

func (m *myRepo) DeleteTodo(ctx context.Context, id int) error {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := m.db.ExecContext(queryCtx, fmt.Sprintf("DELETE FROM todo_item WHERE Id=%d", id))
	if err != nil {
		return err
	}

	return nil
}

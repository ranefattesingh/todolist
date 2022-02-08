package core

import (
	"encoding/json"
	"io"
)

type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type TodoItems []*TodoItem

func (ts *TodoItems) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(ts)
}

func (ts *TodoItem) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(ts)
}

func (t *TodoItem) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(t)
}

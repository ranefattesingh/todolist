package api

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ranefattesingh/todolist/core"
)

type TodoHandler interface {
	GetAll(rw http.ResponseWriter, r *http.Request)
	UpdateTodo(rw http.ResponseWriter, r *http.Request)
	AddTodo(rw http.ResponseWriter, r *http.Request)
	UpdateStatus(rw http.ResponseWriter, r *http.Request)
	DeleteTodo(rw http.ResponseWriter, r *http.Request)
}

type handler struct {
	s core.TodoService
}

func (h *handler) GetAll(rw http.ResponseWriter, r *http.Request) {
	result, err := h.s.GetAll(r.Context())
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}

	err = result.ToJSON(rw)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *handler) UpdateTodo(rw http.ResponseWriter, r *http.Request) {
	regx, err := regexp.Compile("[0-9]+")
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	idString := regx.FindString(r.URL.Path)
	var id int
	id, err = strconv.Atoi(idString)

	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	todoItem := &core.TodoItem{}
	err = todoItem.FromJSON(r.Body)

	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.s.UpdateTodo(r.Context(), id, todoItem)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *handler) AddTodo(rw http.ResponseWriter, r *http.Request) {
	todoItem := &core.TodoItem{}
	err := todoItem.FromJSON(r.Body)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.s.AddTodo(r.Context(), todoItem)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *handler) UpdateStatus(rw http.ResponseWriter, r *http.Request) {
	regx, err := regexp.Compile("[0-9]+")
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	idString := regx.FindString(r.URL.Path)
	var id int
	id, err = strconv.Atoi(idString)

	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.s.UpdateStatus(r.Context(), id)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *handler) DeleteTodo(rw http.ResponseWriter, r *http.Request) {
	regx, err := regexp.Compile("[0-9]+")
	if err != nil {
		fmt.Println(err)
		return
	}
	idString := regx.FindString(r.URL.Path)
	var id int
	id, err = strconv.Atoi(idString)

	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	err = h.s.DeleteTodo(ctx, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func NewTodoHandler(s core.TodoService) TodoHandler {
	return &handler{s}
}

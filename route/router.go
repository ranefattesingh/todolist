package route

import (
	"net/http"

	"github.com/ranefattesingh/todolist/api"
)

type router struct {
	s api.TodoHandler
}

func NewRouter(s api.TodoHandler) *router {
	return &router{s}
}

func (rt *router) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rt.s.GetAll(rw, r)
	case http.MethodPost:
		rt.s.AddTodo(rw, r)
	case http.MethodDelete:
		rt.s.DeleteTodo(rw, r)
	case http.MethodPut:
		rt.s.UpdateTodo(rw, r)
	case http.MethodPatch:
		rt.s.UpdateStatus(rw, r)
	}
}

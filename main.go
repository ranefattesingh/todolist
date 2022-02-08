package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ranefattesingh/todolist/api"
	"github.com/ranefattesingh/todolist/core"
	"github.com/ranefattesingh/todolist/repo/psql"
	"github.com/ranefattesingh/todolist/route"
)

var handler api.TodoHandler

func main() {
	var repo core.TodoRepo
	var err error
	repo, err = psql.NewRepo()

	if err != nil {
		fmt.Println(err)
		return
	}

	var service core.TodoService = core.NewTodoService(repo)
	apiMethods := api.NewTodoHandler(service)
	router := route.NewRouter(apiMethods)

	servMux := http.NewServeMux()
	servMux.Handle("/", router)

	s := http.Server{
		Addr:    ":8000",
		Handler: servMux,
	}

	// TODO: find out why?
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	// Graceful shutdown
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	sig := <-ch
	fmt.Println("System will shoudown in t - 30 seconds", sig)
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ranefattesingh/todolist/core"
	"github.com/ranefattesingh/todolist/repo/psql"
)

func main() {
	var repo core.TodoRepo
	var err error
	repo, err = psql.NewRepo()

	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			var ctx = context.Background()
			result, err := repo.GetAll(ctx)
			if err != nil {
				fmt.Println(err)
			}

			r, err := json.Marshal(result)
			if err != nil {
				fmt.Println(err)
			}
			if err != nil {
				fmt.Println(err)
				return
			}
			rw.Write(r)
		} else if r.Method == http.MethodPut {
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
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			todoItem := &core.TodoItem{}
			err = json.Unmarshal(b, todoItem)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = repo.UpdateTodo(ctx, id, todoItem)
			if err != nil {
				fmt.Println(err)
				return
			}

			rw.WriteHeader(http.StatusOK)
		} else if r.Method == http.MethodPost {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			todoItem := &core.TodoItem{}
			err = json.Unmarshal(b, todoItem)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = repo.AddTodo(r.Context(), todoItem)

			switch r.Context().Err() {
			case context.Canceled:
				fmt.Println("Request is cancelled by user.")
				rw.WriteHeader(http.StatusNotModified)
				return

			case context.DeadlineExceeded:
				fmt.Println("Request Timeout.")
				rw.WriteHeader(http.StatusRequestTimeout)
				return
			}

			if err != nil {
				fmt.Println(err)
				return
			}
		} else if r.Method == http.MethodDelete {
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
			err = repo.DeleteTodo(ctx, id)
			if err != nil {
				fmt.Println(err)
				return
			}
			rw.WriteHeader(http.StatusOK)
		} else if r.Method == http.MethodPatch {
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
			err = repo.UpdateStatus(ctx, id)
			if err != nil {
				fmt.Println(err)
				return
			}
			rw.WriteHeader(http.StatusOK)
		}

	})
	http.ListenAndServe(":8000", nil)
}

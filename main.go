package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ranefattesingh/todolist/repo/psql"
)

func main() {
	repo, err := psql.NewRepo()

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
			// regx, err := regexp.Compile("[0-9]+")
			// if err != nil {
			// 	fmt.Println(err)
			// }
			// idString := regx.FindString(r.URL.Path)
			// var id int
			// id, err = strconv.Atoi(idString)
			// var todos todoItems
			// for i, d := range todos {
			// 	if d.ID == id {
			// 		var b []byte
			// 		b, err = ioutil.ReadAll(r.Body)
			// 		var todo *todoItem = &todoItem{}
			// 		err = json.Unmarshal(b, todo)
			// 		if err != nil {
			// 			fmt.Println(err)
			// 			return
			// 		}
			// 		todo.ID = id
			// 		todos[i] = todo
			// 		fmt.Println(*todos[i])
			// 		return
			// 	}
			// }
			// rw.Write([]byte("Not Found"))
		} else if r.Method == http.MethodPost {
			// b, err := ioutil.ReadAll(r.Body)
			// if err != nil {
			// 	fmt.Println(err)
			// 	return
			// }

			// var todo *todoItem = &todoItem{}
			// json.Unmarshal(b, todo)
			// // todoItems = append(todoItems, todo)
			// rw.WriteHeader(http.StatusCreated)

		} else if r.Method == http.MethodDelete {
			// regx, err := regexp.Compile("[0-9]+")
			// if err != nil {
			// 	fmt.Println(err)
			// }
			// idString := regx.FindString(r.URL.Path)
			// var id int
			// id, err = strconv.Atoi(idString)
			// for i, d := range todoItems {
			// 	if d.ID == id {
			// 		todoItems = append(todoItems[0:i], todoItems[i+1:]...)
			// 		return
			// 	}
			// }
			rw.Write([]byte("Not Found"))
		}
		// rw.Write([]byte("Not Found"))

	})
	http.ListenAndServe(":8000", nil)
}

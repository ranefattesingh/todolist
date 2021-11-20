package main

import "fmt"

type item struct {
	id int
	title string
	description string
	done bool
}

func create() (string, string) {
	var title, description string
	fmt.Print("Enter title for todo:" )
	fmt.Scanln(&title)
	fmt.Print("Enter description for todo:" )
	fmt.Scanln(&description)
	return title, description
}

func update() (string, string) {
	var title, description string
	fmt.Print("Enter title for todo:" )
	fmt.Scanln(&title)
	fmt.Print("Enter description for todo:" )
	fmt.Scanln(&description)
	return title, description
}

func find(items []item, title string) int {
	for i := 0; i < len(items); i++ {
		if items[i].title == title {
			return i
		}
	}
	return -1
}

func main() {
	var (
		choice 	int
		items	[]item
	)

	for {
		fmt.Println("Select anyone: ")
		fmt.Println("1. Add new todo iteam ")
		fmt.Println("2. Update todo item ")
		fmt.Println("3. Delete todo item ")
		fmt.Println("4. Mark/unmark todo item ")
		fmt.Println("5. Find a specific todo item ")
		fmt.Println("6. Display all todo items ")
		fmt.Println("Press any other key to exit")
		fmt.Scanln(&choice)
		if (choice <= 0 || choice > 6) {
			break
		}

		switch choice {
			case 1: 
				t, d := create()
				index := find(items, t)
				if index == -1 {
					newTodoItem := item{len(items) + 1, t, d, false}
					items = append(items, newTodoItem)
				} else {
					fmt.Printf("todo item with %s already exist.", t)
				}


			case 2: 
				var oldTitle string
				fmt.Print("Enter todo title you wish to update: ")
				fmt.Scanln(&oldTitle)
				index := find(items, oldTitle)
				if index >= 0 {
					t, d := update()
					index2 := find(items, t)
					if index2 == -1 || index2 == index {
						items[index].title = t
						items[index].description = d
						fmt.Println("Item is updated.")
					} else {
						fmt.Printf("todo item with %s already exist.", t)
					}
				} else {
					fmt.Println("todo item dosen't exist.")
				}

			case 3: 
				var title string
				fmt.Print("Enter todo title you wish to delete: ")
				fmt.Scanln(&title)
				index := find(items, title)
				if index >= 0 {
					if !items[index].done {
						fmt.Print("todo you're trying to delete is marked as not done proceed? (Y/N)")
						var c string
						fmt.Scanln(&c)
						switch c {
						case "Y":
							fallthrough
						case "y":
							items = append(items[:index], items[index + 1: ]...)
							fmt.Println("Item is deleted.")
						default:
							fmt.Println("todo item is not deleted.")
						}
					}

				} else {
					fmt.Println("todo item dosen't exist.")
				}

			case 4: 
				var title string
				fmt.Print("Enter todo title you wish to mark unmark done: ")
				fmt.Scanln(&title)
				index := find(items, title)
				if index >= 0 {
					items[index].done = !items[index].done
					fmt.Println("Item is updated.")
				} else {
					fmt.Println("todo item dosen't exist.")
				}

			case 5: 
				var title string
				fmt.Print("Enter todo title you wish to find: ")
				fmt.Scanln(&title)
				index := find(items, title)
				if index >= 0 {
					fmt.Println(title)
				} else {
					fmt.Println("todo item dosen't exist.")
				}

			case 6:
				for i := 0; i < len(items); i++ {
					fmt.Println(items[i])
				}
		}
	}
}
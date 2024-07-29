package main

import (
	"fmt"

	"github.com/thediggu/godo/internal/menu"
)

const FILE_NAME = "todo.todofile"

// 1. Take continuouse input until user quits - done
// 2. Marks todos as done
// 3. Delete todos

func main() {
	menu.ClearScreen()
	var user_input string

	for user_input != "4" {
		menu.PrintMenu()
		user_input, user_input_err := menu.TakeInput()

		if user_input_err != nil {
			fmt.Println("Whoops something went wrong")
		} else {
			menu.PerformAction(user_input)
		}
	}
}

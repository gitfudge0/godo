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
	menu.PrintMenu()
	var user_input string

	for user_input != "4" {
		var menu_input string
		fmt.Scan(&menu_input)
		menu.PerformAction(menu_input)
	}
}

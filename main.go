package main

import (
	"fmt"
	"os"

	"github.com/gitfudge0/godo/internal/menu"
)

func OpenFile() (*os.File, error) {
	const FILE_NAME = "todo.todofile"
	_, fileStatErr := os.Stat(FILE_NAME)

	if fileStatErr != nil {
		file, err := os.Create(FILE_NAME)
		return file, err
	}

	file, err := os.OpenFile(FILE_NAME, os.O_CREATE|os.O_RDWR, os.ModeAppend)
	return file, err
}

func main() {
	_, fileErr := OpenFile()

	if fileErr != nil {
		fmt.Println(fmt.Errorf("error opening todo file: %w", fileErr))
	}

	menu.ClearScreen()
	menu.PrintMenu()
	var user_input string

	for user_input != "4" {
		var menu_input string
		fmt.Scan(&menu_input)
		menu.PerformAction(menu_input)
	}
}

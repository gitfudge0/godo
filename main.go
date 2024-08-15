package main

import (
	"fmt"

	"github.com/gitfudge0/godo/internal/fileops"
	"github.com/gitfudge0/godo/internal/menu"
)

func main() {
	file, fileErr := fileops.OpenFile()

	if fileErr != nil {
		fmt.Println(fmt.Errorf("error opening todo file: %w", fileErr))
	}
	defer file.Close()

	todoFile := fileops.TodoFile{
		File: file,
	}

	menu := menu.Menu{
		TodoFile: &todoFile,
	}
	menu.TodoFile.ParseFile()

	var input string

	for input != "4" {
		menu.ClearScreen()
		menu.PrintItems()
		menu.PrintMenu()
		input, _ := menu.TakeInput()
		menu.PerformAction(input)
	}
}

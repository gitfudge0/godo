package fileops

import (
	"fmt"
	"os"
)

func WriteTodo(file *os.File, todoJsonString string) {
	_, writeErr := file.WriteString(todoJsonString)
	if writeErr != nil {
		fmt.Println("Whoops")
		fmt.Println(writeErr)
		os.Exit(1)
	}
}

package fileops

import (
	"os"
)

func OpenFile() (*os.File, error) {
	const FILE_NAME = "todo.todofile"
	_, fileStatErr := os.Stat(FILE_NAME)

	if fileStatErr != nil {
		file, err := os.Create(FILE_NAME)
		return file, err
	}

	file, err := os.OpenFile(FILE_NAME, os.O_RDWR, os.ModeAppend)
	return file, err
}

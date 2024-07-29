package fileops

import (
	"os"
)

func OpenFile(FILE_NAME string) (*os.File, error) {
	_, fileStatErr := os.Stat(FILE_NAME)

	if fileStatErr != nil {
		file, err := os.Create(FILE_NAME)
		return file, err
	}

	file, err := os.OpenFile(FILE_NAME, os.O_RDWR, os.ModeAppend)
	return file, err
}

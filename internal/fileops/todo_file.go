package fileops

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/gitfudge0/godo/internal/models"
)

type TodoFile struct {
	File          *os.File
	parsedContent models.TodoList
}

type TodoFileInterface interface {
	updateContent(list models.TodoList)
	triggerWrite()
	GetFileContent() models.TodoList
	ParseFile() (string, error)
	writeFile(content string) error

	AddTodoItem(models.TodoItem)
	ToggleStatus(int)
	DeleteItem(int)
}

func (file *TodoFile) GetFileContent() models.TodoList {
	return file.parsedContent
}

func (todoFile *TodoFile) ParseFile() error {
	scanner := bufio.NewScanner(todoFile.File)
	for scanner.Scan() {
		text := scanner.Text()
		todoItem := models.TodoItem{}
		json.Unmarshal([]byte(text), &todoItem)
		todoFile.AddTodoItem(todoItem)
	}

	return nil
}

func (todoFile *TodoFile) updateFileContent(list models.TodoList) {
	todoFile.parsedContent = list
	todoFile.triggerWrite()
}

func (todoFile *TodoFile) triggerWrite() {
	todoFile.writeFile(todoFile.parsedContent.ToJsonString())
}

func (todoFile *TodoFile) writeFile(content string) {
	todoFile.File.Truncate(0)
	todoFile.File.Seek(0, 0)
	todoFile.File.WriteString(content)
}

func (todoFile *TodoFile) AddTodoItem(item models.TodoItem) {
	result := todoFile.parsedContent.AddItem(item)
	todoFile.updateFileContent(result)
}

func (todoFile *TodoFile) ToggleStatus(index int) {
	result := todoFile.parsedContent.ToggleStatus(index)
	todoFile.updateFileContent(result)
}

func (todoFile *TodoFile) DeleteItem(index int) {
	result := todoFile.parsedContent.RemoveByIndex(index)
	todoFile.updateFileContent(result)
}

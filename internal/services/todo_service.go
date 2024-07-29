package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/thediggu/godo/internal/enum"
	"github.com/thediggu/godo/internal/fileops"
	"github.com/thediggu/godo/internal/models"
)

const file_name = "todo.todofile"

func getParsedTodoItem(text string) models.TodoItem {
	todo := models.TodoItem{}
	json.Unmarshal([]byte(text), &todo)
	return todo
}

func IsPriorityValid(priority string) bool {
	pslice := []string{enum.HighPriority.PriorityString(), enum.MediumPriority.PriorityString(), enum.LowPriority.PriorityString()}
	return slices.Contains(pslice, priority)
}

func determineNewIndex(file *os.File) (int, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	if size := fileInfo.Size(); size == 0 {
		return 1, nil
	}

	scanner := bufio.NewScanner(file)

	var text string
	for scanner.Scan() {
		text = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	lastEntry := getParsedTodoItem(text)
	return lastEntry.Id + 1, nil
}

func readTodos(file *os.File) (todoList []models.TodoItem, tlerr error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if size := fileInfo.Size(); size == 0 {
		return nil, nil
	}

	scanner := bufio.NewScanner(file)

	var text string
	for scanner.Scan() {
		text = scanner.Text()
		todoItem := getParsedTodoItem(text)
		todoList = append(todoList, todoItem)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return todoList, nil
}

func ListTodos() {
	file, file_err := fileops.OpenFile(file_name)
	if file_err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	// Read each todo from file
	todoList, tlerr := readTodos(file)
	if tlerr != nil {
		fmt.Println("Error reading todos")
		return
	}

	if len(todoList) == 0 {
		fmt.Println("Add some todo items to view them")
		return
	}

	fmt.Println("Your todos:")
	for i, v := range todoList {
		checkmark := "[ ]"

		if v.Is_done {
			checkmark = "[X]"
		}

		var indexStr string

		if i < 9 {
			indexStr = fmt.Sprintf("0%d", i+1)
		} else {
			indexStr = fmt.Sprintf("%d", i+1)
		}

		str := fmt.Sprintf("%s. %s %s", indexStr, checkmark, v.Title)
		fmt.Println(str)
	}
}

func AddTodo(todoInput, priority string) {
	if !IsPriorityValid(priority) {
		fmt.Println("Invalid priority")
		os.Exit(1)
	}

	file, file_err := fileops.OpenFile(file_name)
	if file_err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	id, idErr := determineNewIndex(file)
	if idErr != nil {
		fmt.Println("Whoops")
		fmt.Println(idErr)
		return
	}

	todoItem := models.TodoItem{
		Created_at: time.Now(),
		Title:      todoInput,
		Is_done:    false,
		Priority:   priority,
		Id:         id,
	}

	fileops.WriteTodo(file, todoItem.JsonString())
}

func DeleteTodo(id int) {
	file, file_err := fileops.OpenFile(file_name)
	if file_err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	lineNum := 1

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if lineNum == id {
			// OK, I'll figure this out later - gonna go eat dinner. Ciao!
		}
		_ = scanner.Text()
		lineNum += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error deleting item")
		return
	}
}

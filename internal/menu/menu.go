package menu

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gitfudge0/godo/internal/fileops"
	"github.com/gitfudge0/godo/internal/models"
)

func determineNewIndex(file *fileops.TodoFile) int {
	contentLen := len(file.GetFileContent())
	if contentLen == 0 {
		return 1
	}
	return contentLen + 1
}

func waitForEnter() {
	fmt.Println("\nPress enter to continue")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func isOptionValid(input string) bool {
	validOptions := []string{"1", "2", "3", "4"}
	return slices.Contains(validOptions, input)
}

func handleError(err error) {
	errString := fmt.Errorf("there was an error: %w", err)
	fmt.Println(errString)
	os.Exit(2)
}

func isPriorityValid(priority string) bool {
	pslice := []string{"high", "medium", "low"}
	return slices.Contains(pslice, priority)
}

type Menu struct {
	TodoFile *fileops.TodoFile
}

type MenuInterface interface {
	ClearScreen()
	PrintMenu()
	PrintItems()
	TakeInput() (string, error)
	PerformAction(string) error

	doAddItem()
	doToggleStatus()
	doDeleteItem()
}

func (menu *Menu) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (menu *Menu) PrintMenu() {
	fmt.Println("Enter your choice: ")
	fmt.Println("1. Add a todo")
	fmt.Println("2. Toggle status")
	fmt.Println("3. Delete a todo")
	fmt.Println("4. Exit")
}

func (menu *Menu) PrintItems() {
	for i, v := range menu.TodoFile.GetFileContent() {
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

	fmt.Println("")
}

func (menu *Menu) TakeInput() (input string, err error) {
	reader := bufio.NewReader(os.Stdin)
	user_input, _ := reader.ReadString('\n')
	user_input = strings.TrimSpace(user_input)

	return user_input, nil
}

func (menu *Menu) PerformAction(menuInput string) error {
	iv := isOptionValid(menuInput)
	if !iv {
		fmt.Println("Oof")
	}

	menu.ClearScreen()
	switch menuInput {
	case "1":
		menu.doAddItem()
	case "2":
		menu.doToggleStatus()
	case "3":
		menu.doDeleteItem()
	case "4":
		os.Exit(0)
	}
	return nil
}

func (menu *Menu) doAddItem() {
	index := determineNewIndex(menu.TodoFile)

	fmt.Println("Enter todo")
	todoInput, _ := menu.TakeInput()
	menu.ClearScreen()

	fmt.Println("Enter priority(high/medium/low)")
	priorityInput, _ := menu.TakeInput()
	menu.ClearScreen()

	for !isPriorityValid(strings.ToLower(priorityInput)) {
		fmt.Println("Invalid priority", priorityInput)
		waitForEnter()
	}

	item := models.TodoItem{
		Id:         index,
		Title:      todoInput,
		Priority:   priorityInput,
		Created_at: time.Now(),
		Is_done:    false,
	}

	menu.TodoFile.AddTodoItem(item)
}

func (menu *Menu) doToggleStatus() {
	menu.PrintItems()
	fmt.Println("Enter todo item # to toggle")
	todoInput, _ := menu.TakeInput()

	index, err := strconv.Atoi(todoInput)
	if err != nil {
		handleError(err)
	}

	menu.TodoFile.ToggleStatus(index)
}

func (menu *Menu) doDeleteItem() {
	menu.PrintItems()
	fmt.Println("Enter todo item # to delete")
	todoInput, _ := menu.TakeInput()

	index, err := strconv.Atoi(todoInput)
	if err != nil {
		handleError(err)
	}

	menu.TodoFile.DeleteItem(index)
}

package menu

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"github.com/thediggu/godo/internal/services"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func PrintMenu() {
	fmt.Println("Enter your choice: ")
	fmt.Println("1. List all todos")
	fmt.Println("2. Add a todo")
	fmt.Println("3. Delete a todo")
	fmt.Println("4. Exit")
}

func getChoices() []string {
	return []string{"1", "2", "3", "4"}
}

func ValidateMenuInput(choice string) bool {
	return slices.Contains(getChoices(), choice)
}

func TakeInput() (input string, err error) {
	reader := bufio.NewReader(os.Stdin)
	user_input, _ := reader.ReadString('\n')
	user_input = strings.TrimSpace(user_input)

	return user_input, nil
}

func insertTodoEntry() {
	fmt.Println("Enter todo: ")
	todoInput, tierr := TakeInput()
	if tierr != nil {
		fmt.Println("Whoops")
		fmt.Println(tierr)
		return
	}

	fmt.Println("Enter priority(high/medium/low): ")
	priorityInput, perr := TakeInput()
	if perr != nil {
		fmt.Println("Whoops")
		fmt.Println(tierr)
		return
	}

	services.AddTodo(todoInput, priorityInput)
	ClearScreen()
	services.ListTodos()
}

func waitForEnter() {
	fmt.Println("\nPress enter to continue")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func preAction() {
	ClearScreen()
}

func postAction() {
	waitForEnter()
	ClearScreen()
}

func PerformAction(choice string) {
	preAction()
	switch choice {
	case "1":
		services.ListTodos()
	case "2":
		insertTodoEntry()
	case "3":
		// This was supposed to do something
	case "4":
		fmt.Println("Bye!")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
	postAction()
}

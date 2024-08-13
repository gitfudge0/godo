package menu

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/gitfudge0/godo/internal/services"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func PrintMenu() {
	services.ListTodos()
	fmt.Println("")
	fmt.Println("Enter your choice: ")
	fmt.Println("1. Toggle status")
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
	ClearScreen()

	var priorityInput string
	var perr error
	for !services.IsPriorityValid(strings.ToLower(priorityInput)) {
		fmt.Println("Enter priority(high/medium/low): ")
		priorityInput, perr = TakeInput()

		if perr != nil {
			fmt.Println("Whoops")
			fmt.Println(tierr)
			return
		}

		if !services.IsPriorityValid(strings.ToLower(priorityInput)) {
			fmt.Println("Invalid priority", priorityInput)
			waitForEnter()
		}
		ClearScreen()
	}

	services.AddTodo(todoInput, strings.ToLower(priorityInput))
	ClearScreen()
	services.ListTodos()
}

func initiateDelete() {
	services.ListTodos()
	fmt.Println("\nEnter ID to delete")
	var input string
	fmt.Scan(&input)
	fmt.Println(input)
	if i, err := strconv.Atoi(input); i > 0 && err == nil {
		services.DeleteTodo(i)
	}
	ClearScreen()
	PrintMenu()
}

func waitForEnter() {
	fmt.Println("\nPress enter to continue")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func waitForToggleInput() {
	services.ListTodos()
	fmt.Println("\nEnter the index of todo to toggle status or press enter to continue")
	var input string
	fmt.Scan(&input)
	fmt.Println(input)
	if i, err := strconv.Atoi(input); i > 0 && err == nil {
		services.ToggleTodoStatus(i)
	}
	ClearScreen()
	PrintMenu()
}

func preAction() {
	ClearScreen()
}

func PerformAction(choice string) {
	preAction()
	switch choice {
	case "1":
		waitForToggleInput()
	case "2":
		insertTodoEntry()
		waitForEnter()
		ClearScreen()
		PrintMenu()
	case "3":
		initiateDelete()
	case "4":
		fmt.Println("Bye!")
		os.Exit(0)
	default:
		ClearScreen()
		PrintMenu()
	}
}

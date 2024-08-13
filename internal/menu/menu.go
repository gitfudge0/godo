package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gitfudge0/godo/internal/fileops"
)

type Menu struct {
	todoFile *fileops.TodoFile
}

type MenuInterface interface {
	ClearScreen()
	PrintMenu()
	TakeInput() (string, error)
}

func (menu *Menu) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (menu *Menu) PrintMenu() {
	fmt.Println("")
	fmt.Println("Enter your choice: ")
	fmt.Println("1. Toggle status")
	fmt.Println("2. Add a todo")
	fmt.Println("3. Delete a todo")
	fmt.Println("4. Exit")
}

func (menu *Menu) TakeInput() (input string, err error) {
	reader := bufio.NewReader(os.Stdin)
	user_input, _ := reader.ReadString('\n')
	user_input = strings.TrimSpace(user_input)

	return user_input, nil
}

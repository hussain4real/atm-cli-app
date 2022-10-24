package main

import "fmt"

var menuNumber int

func main() {
	fmt.Println("hello world")
}
func mainMenu() {
	newLine(1)
	fmt.Println("Select operation")
	fmt.Println("1. Create Todo \t")
	fmt.Println("2. List Todo Item \t")
	fmt.Println("3. Edit Todo Item \t")
	fmt.Println("4. Delete Todo item")
	fmt.Println("0. Exit the program")
	_, err := fmt.Scan(&menuNumber)

	if err != nil {
		fmt.Println("Error:, please select the correct menu item")

	}

	switch menuNumber {
	case 1:
		createTodo()
	case 2:
		listTodo()
	case 3:
		updateTodo()
	case 4:
		deleteTodo()
	case 0:
		exitProgram()
	default:
		fmt.Println("Error: Invalid input")
	}
}

func newLine(numberOfLines int) {
	for i := 0; i < numberOfLines; {
		fmt.Println("\n")
		i++
	}
}

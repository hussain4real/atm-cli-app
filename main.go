package main

import (
	"fmt"
	"os"
)

var menuNumber int
var pin = 0000
var accountBalance = 1000.00

func main() {
	welcome()
	//create a prompt for the user to enter their pin
	fmt.Println("Please enter your pin")
	//scan the pin
	_, err := fmt.Scan(&pin)
	if err != nil {
		return
	}
	//check if the pin is correct
	if pin == 0000 {
		//if the pin is correct, run the main menu
		mainMenu()
	} else {
		//if the pin is incorrect, print an error message
		fmt.Println("Error: Incorrect pin")
		//run the main function again
		main()
		//if the user enters the pin incorrectly 3 times, display an error message and exit the program

	}

}

func welcome() {
	fmt.Println("ππΎππΎππΎππΎ Welcome to Hussain atm CLI app where banking is flawless. ππΎππΎππΎππΎ")
	newLine(1)
}

func mainMenu() {
	newLine(1)
	fmt.Println("Select operation:")
	fmt.Println("1. Change pin \t")
	fmt.Println("2. Check account balance \t")
	fmt.Println("3. Withdraw money \t")
	fmt.Println("4. Deposit money")
	fmt.Println("0. Cancel/Exit the program")
	_, err := fmt.Scan(&menuNumber)

	if err != nil {
		fmt.Println("Error:, please select the correct menu item")

	}

	switch menuNumber {
	case 1:
		changePin()
	case 2:
		checkAccountBalance()
	case 3:
		withdrawMoney()
	case 4:
		DepositMoney()
	case 0:
		exitProgram()
	default:
		fmt.Println("Error: Invalid input")
	}
}

func changePin() {
	newLine(1)
	fmt.Println("Please enter your new pin π:")
	var newPin int
	_, err := fmt.Scan(&newPin)
	if err != nil {
		fmt.Println("Error: Invalid input β½")
	}
	pin = newPin
	fmt.Println("Pin changed successfully π")
	performAnotherOperation()
}

// check account balance
func checkAccountBalance() {
	newLine(1)
	fmt.Println("Your account balance is: ", accountBalance)
	performAnotherOperation()
}

// withdraw money
func withdrawMoney() {
	newLine(1)
	fmt.Println("Enter the amount you want to withdraw π΅:")
	var amount float64
	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("Error: Invalid input β½")
	}
	if amount > accountBalance {
		fmt.Println("Error: Insufficient funds π₯Ή")
	} else {
		accountBalance -= amount
		fmt.Println("Withdrawal successful π€")
	}
	performAnotherOperation()
}

// DepositMoney deposit money
func DepositMoney() {
	newLine(1)
	fmt.Println("Enter the amount you want to deposit:")
	var amount float64
	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("Error: Invalid input")
	}
	accountBalance += amount
	fmt.Println("Deposit successful π€")
	performAnotherOperation()
}

func performAnotherOperation() {
	fmt.Println("Do you want to perform another operation? (Y/N)")
	var answer string
	_, err := fmt.Scan(&answer)
	if err != nil {
		fmt.Println("Error: Invalid input")
	}
	if answer == "Y" {
		mainMenu()
	} else if answer == "N" {
		exitProgram()
	} else {
		fmt.Println("Error: Invalid input")
	}
}

// exit program
func exitProgram() {
	fmt.Println("Thanks for banking with us π€")
	os.Exit(0)
}

func newLine(numberOfLines int) {
	for i := 0; i < numberOfLines; {
		fmt.Println("\n")
		i++
	}
}

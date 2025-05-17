package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000.00 // for better code readability declare the variable with a type

	fmt.Println("Welcome to the bank!")
	fmt.Println("What would like to do today?")

	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var choice int
	fmt.Println("Please enter your choice (1-4):")
	fmt.Scanln(&choice)
	fmt.Println("You chose option: ", choice)

	// wantsCheckBalance := choice == 1
	// if wantsCheckBalance {
	if choice == 1 {
		fmt.Println("Your balance is:", accountBalance)
	} else if choice == 2 {
		fmt.Println("Enter amount to deposit:")
		var depositAmount float64
		fmt.Scanln(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid deposit amount!")
			return
		}

		accountBalance += depositAmount
		fmt.Println("You have deposited:", depositAmount)
		fmt.Println("Your new balance is:", accountBalance)
	} else if choice == 3 {
		fmt.Println("Enter amount to withdraw:")
		var withdrawAmount float64
		fmt.Scanln(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Println("Invalid withdraw amount!")
			return
		}

		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient funds!")
			return
		} else {
			accountBalance -= withdrawAmount
			fmt.Println("You have withdrawn:", withdrawAmount)
			fmt.Println("Your new balance is:", accountBalance)
		}
	} else {
		fmt.Println("Goodbye!")
	}
}

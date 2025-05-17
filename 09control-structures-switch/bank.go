package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000.00 // for better code readability declare the variable with a type

	fmt.Println("Welcome to the bank!")

	// for i := 0; i < 2; i++ {
	// infinite loop
	for {

		fmt.Println("What would like to do today?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Println("Please enter your choice (1-4):")
		fmt.Scanln(&choice)
		fmt.Println("You chose option: ", choice)

		// The switch in go does not require a break statement at the end of each case!
		// If specified it just breaks out from the switch statement
		// If you need to break out of a loop then switch statement is not the right choice.
		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Println("Enter amount to deposit:")
			var depositAmount float64
			fmt.Scanln(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount!")
				// return // no longer needed as we are using a loop and exiting on choice 4
				continue // continue to the next iteration of the loop
			}

			accountBalance += depositAmount
			fmt.Println("You have deposited:", depositAmount)
			fmt.Println("Your new balance is:", accountBalance)
		case 3:
			fmt.Println("Enter amount to withdraw:")
			var withdrawAmount float64
			fmt.Scanln(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid withdraw amount!")
				// return // no longer needed as we are using a loop and exiting on choice 4
				continue // continue to the next iteration of the loop
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds!")
				// return // no longer needed as we are using a loop and exiting on choice 4
				continue // continue to the next iteration of the loop
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("You have withdrawn:", withdrawAmount)
				fmt.Println("Your new balance is:", accountBalance)
			}
		case 4:
			fmt.Println("Goodbye!")
			// break // would exit switch, but no the for loop
			return // exit the program
		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}

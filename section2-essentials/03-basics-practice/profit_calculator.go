package main

import "fmt"

// Practice basics like variable declaration and usage, user input, output to console, and basic math operations.
func main() {
	// define variables as float64, so we can not have to convert later and user can input decimal values
	var revenue float64
	var expenses float64
	var taxRate float64

	// fmt.Print("Enter the revenue, expense and tax rate: ")
	// fmt.Scan(&revenue, &expenses, &taxRate) // get values into multiple variable at once, space or enter is a delimiter when receiving input
	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println("Earnings Before Tax: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Profit Ratio: ", ratio)
}

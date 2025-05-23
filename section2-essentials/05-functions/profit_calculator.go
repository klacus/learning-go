package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Enter the revenue: ")
	// fmt.Scan(&revenue)
	revenue := getUserInput("Enter the revenue: ")

	// fmt.Print("Enter the expenses: ")
	// fmt.Scan(&expenses)
	expenses := getUserInput("Enter the expenses: ")

	// fmt.Print("Enter the tax rate: ")
	// fmt.Scan(&taxRate)
	taxRate := getUserInput("Enter the tax rate: ")

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit
	ebt, profit, ratio := calculateFinancialMetrics(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Profit Ratio: %.3f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancialMetrics(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

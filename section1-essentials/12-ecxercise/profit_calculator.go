package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Enter the revenue: ")
	// fmt.Scan(&revenue)
	revenue, err := getUserInput("Enter the revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Print("Enter the expenses: ")
	// fmt.Scan(&expenses)
	expenses, err := getUserInput("Enter the expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Print("Enter the tax rate: ")
	// fmt.Scan(&taxRate)
	taxRate, err := getUserInput("Enter the tax rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// if err1 != nil || err2 != nil || err3 != nil {
	// 	fmt.Println("Error: One or more inputs are invalid.")
	// 	fmt.Println(err1)
	// 	fmt.Println(err12
	// 	fmt.Println(err3
	// 	return
	// }

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit
	ebt, profit, ratio := calculateFinancialMetrics(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Profit Ratio: %.3f\n", ratio)
	storeResuts(ebt, profit, ratio)
}

func storeResuts(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.3f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
	fmt.Println("Results stored successfully.")
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	// Check if the user input is negative
	if userInput <= 0 {
		fmt.Println("Error: Input must be a positive number.")
		return 0, errors.New("input must be a positive number")
	}

	return userInput, nil
}

func calculateFinancialMetrics(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

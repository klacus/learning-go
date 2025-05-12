package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate float64 = 2.5
	var investmentAmount float64 // declaration only without default value, always must define the type if declared this way
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter the amount you want to invest: ")
	fmt.Scan(&investmentAmount) // user input, using a pointer to update the value of investmentAmount variable, fmt.Scan() can get a single value for every variable defined for input, it can get values for more than one variable at a time
	fmt.Print("Enter the number of years you want to invest: ")
	fmt.Scan(&years)
	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future Value (adjusted for inflation): ", futureRealValue)
	// fmt.Printf("Future Value: %v\nFuture Value (adjusted for inflation): %v", futureValue, futureRealValue) // %v is a default format specifier, it will print the value in its default format
	// fmt.Printf("Future Value: %.0f\nFuture Value (adjusted for inflation): %.0f", futureValue, futureRealValue) // %v is a default format specifier, it will print the value in its default format
	// fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for inflation): %.1f", futureValue, futureRealValue) // %v is a default format specifier, it will print the value in its default format

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV) // print the formatted string

	// If longer text is needed, broken out into multiple lines. Mind the indentation and line breaks, they all matter.
	// fmt.Printf(`Future Value: %.1f
	//Future Value (adjusted for inflation): %.1f`, futureValue, futureRealValue) // %v is a default format specifier, it will print the value in its default format

}

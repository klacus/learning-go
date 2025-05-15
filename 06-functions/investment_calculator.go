package main

import (
	"fmt"
	"math"
)

// You can add constants and variables outside the main or any function. But you can not use shortcut variable declaration (:=) outside a function.
const inflationRate float64 = 2.5

func main() {

	var investmentAmount float64 // declaration only without default value, always must define the type if declared this way
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Enter the amount you want to invest: ")
	outputText("Enter the amount you want to invest: ") // function call

	fmt.Scan(&investmentAmount) // user input, using a pointer to update the value of investmentAmount variable, fmt.Scan() can get a single value for every variable defined for input, it can get values for more than one variable at a time

	fmt.Print("Enter the number of years you want to invest: ")
	outputText("Enter the number of years you want to invest: ") // function call

	fmt.Scan(&years)

	fmt.Print("Enter the expected return rate: ")
	outputText("Enter the expected return rate: ") // function call

	fmt.Scan(&expectedReturnRate)

	// If a function returns multiple values then the number of variables must be equal to the number of return values.
	futureValue, futureRealValue := calculateFutureValue(investmentAmount, years, expectedReturnRate)
	// var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV) // print the formatted string

}

// We can define our own functions.
// The function need a name and optionally some input parameter(s) and optionally a return type.
// func outputText(text string) {
// func outputText(someNumber int, text string) {
func outputText(text string) {
	fmt.Println(text)
}

// The function can return a value or multiple values.
// Just like variable definition, function input parameters can be defined with a type and a name and same type can be used for multiple parameters.
func calculateFutureValue(investmentAmount, years, expectedReturnRate float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

// // Return values can be defined with a type and a name and same type can be used for multiple return values.
// // In this the return variable names already defined for the function.
// func calculateFutureValue(investmentAmount, years, expectedReturnRate float64) (fv float64, rfv float64) {
// 	// If return value names are defined then we can use them directly in the function without defining them again.
// 	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv = fv / math.Pow(1+inflationRate/100, years)

// 	// Also if return value names are defined then we can use the return statement without defining the return values again.
// 	return
// }

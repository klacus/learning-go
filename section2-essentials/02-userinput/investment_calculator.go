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

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

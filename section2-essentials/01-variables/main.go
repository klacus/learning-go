package main

import (
	"fmt"
	"math"
)

// Variables in Go are declared using the var keyword. The type of the variable is specified after the variable name.
// Go supports multiple variable declarations in one line, and the type can be inferred from the value assigned to the variable.
// Go also supports short variable declarations using the := operator, which infers the type from the value assigned to the variable.
// The short variable declaration can only be used inside functions, not at the package level.
func main() {

	const inflationRate float64 = 2.5 // Constant declaration. Constants are immutable and must be assigned a value at the time of declaration.

	var investmentAmount, years = 1000, 10 // Multiple variable declaration in one line with inferred type. Variables may be different types.
	// var investmentAmount, years float64 = 1000, 10 // Multiple variable declaration in one line with explicit type. All variable will be the same type.
	// investmentAmount, years := 1000, 10 // Short variable declaration for multiple variable declaration in one line with types always inferred.

	var expectedReturnRate float64 = 5.5 // Single variable declaration, with explicit type .
	// var expectedReturnRate = 5.5 // Single variable declaration, type inferred.
	// expectedReturnRate := 5.5 // Short single variable declaration, always inferred.

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, float64(years))

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

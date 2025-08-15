package main

import (
	"fmt"
)

// like an alias
type str string

// then can add custom method to it, which is not possible for the built-in types
func (text str) log() {
	fmt.Print(text)
}

func main() {
	// then you must set the custom type explicitly
	var name str = "John Doe"
	// then use the custom method
	name.log()

}

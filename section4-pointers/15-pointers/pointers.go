package main

import "fmt"

func main() {
	age := 32 // regular variable

	// agePointer *int // declaring a pointer to an int
	agePointer := &age // pointer to the variable

	// fmt.Println("Age:", age) // prints the value of age
	fmt.Println("Age:", *agePointer) // prints the address of the pointer, not the value

	// adultYears := getAdultYears(age)
	// adultYears := getAdultYears(age)
	// adultYears := getAdultYears(&age)       // passing the address of age to the function
	adultYears := getAdultYears(agePointer) //  passing a pointer type variable to the function, want a pointer, get a pointer
	fmt.Println("Adult years:", adultYears)

	// alternative
	getAdultYearsAlternative(agePointer)
	fmt.Println("Adult years, alternative calculation:", age)
}

// func getAdultYears(age int) int {
// 	// age is a pointer to an int
// 	return age - 18 // subtract 18
// }

func getAdultYears(age *int) int {
	// age is a pointer to an int
	return *age - 18 // dereference the pointer and subtract 18

	// alternative
	// *age = *age - 18 // modify the value at the pointer address, not returning it
}

// another alternative
func getAdultYearsAlternative(age *int) {
	// this is discouraged as it overwrites the age variable directly, and that can be unexpected, in this case it is recommended to name the function indicating it will update the referenced variable's value.
	*age = *age - 18 // modify the value at the pointer address, not returning it, it will overwrite the value ov age variable value directly
}

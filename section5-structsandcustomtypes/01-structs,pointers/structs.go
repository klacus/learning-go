package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	appUser := user{
		firstName: getUserData("Please enter your first name: "),
		lastName:  getUserData("Please enter your last name: "),
		birthDate: getUserData("Please enter your birthdate (MM/DD/YYYY): "),
		createdAt: time.Now(),
	}

	// alternatively, then the order of the values must match the struct definition
	// appUser2 := user{
	// 	"Firsname",
	// 	"Lastname",
	// 	"bdate",
	// 	time.Now(),
	// }
	//
	//  or empty instance
	// appUser3 := user{}
	//
	// or have null values or default, then the fields will have their default value from the struct declaration (null or otherwise)
	// appUser4 := user{
	// 	firstName: getUserData("Please enter your first name: "),
	// 	birthDate: getUserData("Please enter your birthdate (MM/DD/YYYY): "),
	// 	createdAt: time.Now(),
	// }

	// ... do something awesome with that gathered data!

	// use pointers also works
	outputUserData(&appUser)
}

func outputUserData(u *user) {
	//  technically correct would be dereferencing it, but shortcut is allowed
	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate, (*u).createdAt.Format("2006-01-02 15:04:05"))
	// short cut is allowed for structs as exception...
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt.Format("2006-01-02 15:04:05"))
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

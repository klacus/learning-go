package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// attach a function to the struct, no parameter needed for the method as it has access to the struct fields
// the "(u user)" argument is called "Receiver Argument" a.k.a. "Receiver"
// !!! the Receiver Arguments are passed around like any other parameters i.e. copies !!!
func (u user) outputUserData() {
	//  technically correct would be dereferencing it, but shortcut is allowed
	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate, (*u).createdAt.Format("2006-01-02 15:04:05"))
	// short cut is allowed for structs as exception...
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt.Format("2006-01-02 15:04:05"))
}

// !!! methods that modify the struct you must use a pointer, otherwise you will edit a copy
// dereferencing the pointer is not necessary to modify the original struct instance
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// constructor function, it is a utility function
// this utility function creates a new user instance with the output type "user"
//
//	func newUser(firstName, lastName, birthDate string) user {
//		return user{
//	 we can return a pointer, to prevent a copy to be returned, save memory, etc.
//
// func newUser(firstName, lastName, birthDate string) *user {
// you can add validation, in which case you need to return an error if validation failed
// this allows adding validation logic in a reuseable way
func newUser(firstName, lastName, birthDate string) (*user, error) {
	// can add validation here, so it will be executed every time a new user is created
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all parameters are required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	// appUser := user{
	// 	firstName: getUserData("Please enter your first name: "),
	// 	lastName:  getUserData("Please enter your last name: "),
	// 	birthDate: getUserData("Please enter your birthdate (MM/DD/YYYY): "),
	// 	createdAt: time.Now(),
	// }
	// appUser := newUser(
	// 	getUserData("Please enter your first name: "),
	// 	getUserData("Please enter your last name: "),
	// 	getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	// )
	appUser, err := newUser(
		getUserData("Please enter your first name: "),
		getUserData("Please enter your last name: "),
		getUserData("Please enter your birthdate (MM/DD/YYYY): "),
	)

	// handle error handling
	if err != nil {
		fmt.Println(err)
		return
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

	// then no arguments needed for the method / function either
	// outputUserData(&appUser)
	appUser.outputUserData()

	// use the another method
	appUser.clearUserName()
	appUser.outputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// fmt.Scan(&value)
	// if validation fail then the regular Scan will just wait, so use Scanln to take the Enter as end of input
	fmt.Scanln(&value)
	return value
}

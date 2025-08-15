package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {

	var appUser *user.User

	// calling the constructor with a commonly named "New" method
	appUser, err := user.New(
		getUserData("Please enter your first name: "),
		getUserData("Please enter your last name: "),
		getUserData("Please enter your birthdate (MM/DD/YYYY): "),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserData()
	appUser.ClearUserName()
	appUser.OutputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

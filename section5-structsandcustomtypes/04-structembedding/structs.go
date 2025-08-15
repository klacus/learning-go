package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {

	var appUser *user.User

	appUser, err := user.New(
		getUserData("Please enter your first name: "),
		getUserData("Please enter your last name: "),
		getUserData("Please enter your birthdate (MM/DD/YYYY): "),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin, err2 := user.NewAdmin(
		getUserData("Please enter your email: "),
		getUserData("Please enter your password: "),
		*appUser, // or add any other new user, just get the data for it
	)

	if err2 != nil {
		fmt.Println(err)
		return
	}

	// admin.User.OutputUserData()
	// admin.User.ClearUserName()
	// admin.User.OutputUserData()
	// you can call the methods on the admin struct itself as the User struct is anonymous
	admin.OutputUserData()
	admin.ClearUserName()
	admin.OutputUserData()

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

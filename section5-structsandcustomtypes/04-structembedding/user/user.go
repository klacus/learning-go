package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// embedding the User struct within Admin struct (like classes and inheritance, go does not have classes! )
type Admin struct {
	email    string
	password string
	// you can embed it with a name, explicit embedding
	// user User
	// or with upper case if you want to expose it
	// User User
	//  or an anonymously, which is the most common case
	User
}

func (u User) OutputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt.Format("2006-01-02 15:04:05"))
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// alternatively you can ask for all the fields of a User struct or just expecting a User struct instance anyways
func NewAdmin(email, password string, user User) (*Admin, error) {
	// func NewAdmin(firstName, lastName, birthDate, email, password string) (*Admin, error) {
	// you need all the validation here too if you get the fields for the User struct instead of an instance of the User struct
	// if  email == "" || password == "" || firstName == "" || lastName == "" || birthDate == "" {
	if email == "" || password == "" {
		return nil, errors.New("email and password parameters are required for admin")
	}

	return &Admin{
		email:    email,
		password: password,
		User:     user,
	}, nil
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all parameters are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

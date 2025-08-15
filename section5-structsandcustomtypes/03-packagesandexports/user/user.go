package user

import (
	"errors"
	"fmt"
	"time"
)

// when exported the name should be upper case, applies to all fields too, if field starts with lower case, it will not be exported
// however the fields may not need to be accessible outside the module, so we can control their value intake
// it is desired when we do not want the fields to be directly interacted with (i.e. they are protected to be validated and managed by the logic within the module)
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u User) OutputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt.Format("2006-01-02 15:04:05"))
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// it is common that the constructor is just called "New" as the caller must call it from the package anyways
// func NewUser(firstName, lastName, birthDate string) (*User, error) {
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

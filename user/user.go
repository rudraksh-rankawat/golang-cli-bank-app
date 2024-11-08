package user

import (
	"time"
	"errors"
	"fmt"
)

type User struct {
	firstName string
	lastName  string
	age       int32
	dob       string
	createdAt time.Time
}

type Admin struct {
	User User
	password string
}

func NewAdmin(firstName, lastName, password string) Admin {
	return Admin{
		User: User{
			firstName: firstName,
			lastName: lastName,	
			age: 12,
			dob: "10/10/1993",
			createdAt: time.Now(),
		},
		password: password,
	}
}

func New(firstName string, lastName string, age int32, dob string) (*User, error) {
	
	if firstName == "" || lastName == "" {
		return nil, errors.New("define firstname and last name both while calling newUser constructor func")
	}

	return &User{
		firstName: firstName,
		lastName: lastName,
		age: age,
		dob: dob,
		createdAt: time.Now(),
	}, nil
}

func (u User) DisplayName() string {
	val := fmt.Sprintf("%v %v", u.firstName, u.lastName)

	return val
}

func (u *User) DisplayAge() string {
	val := fmt.Sprintf("%v", u.age)

	return val
}

func (u *User) ClearName() {
	u.firstName = ""
	u.lastName = ""
}

func (u *User) ChangeAge(newAge int32) {
	u.age = newAge
}


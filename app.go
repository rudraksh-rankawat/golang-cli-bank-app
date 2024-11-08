package main

import (
	"fmt"

	"example.com/structy/user"
)

func main() {

	// firstName := getUserData("please enter your first name: ")
	// lastName := getUserData("Please enter your last name: ")
	// dob := getUserData("Please enter your dob: ")

	admin := user.NewAdmin("Karunesh", "Trivedi", "srasawerqwqwe")

	fmt.Printf("Admin Name: %v\n", admin.User.DisplayName())
	fmt.Printf("Admin Age: %v\n", admin.User.DisplayAge())

}

func getUserData(prompt string) string {
	fmt.Printf("%v", prompt)
	var value string
	fmt.Scanln(&value)
	return value
}

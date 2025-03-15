package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("First name: ")
	lastName := getUserData("Last name: ")
	birthDate := getUserData("Birthdate (MM/DD/YYYY): ")

	appUser, error := user.New(firstName, lastName, birthDate)

	if error != nil {
		fmt.Print(error)
		return
	}

	admin := user.NewAdmin("admin@gmail.com", "admin123")
	admin.User.OutputUserDetails()
	admin.OutputUserDetails() // both work!

	appUser.OutputUserDetails()
	appUser.ClearUsername()
	appUser.OutputUserDetails()
}

func getUserData(promtText string) string {
	fmt.Print(promtText)
	var value string
	fmt.Scanln(&value)
	return value
}

package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Phone string
}

func Print_User_Details(usr User) {
	fmt.Printf("Name: %s\nAge: %d\nPhone: %s\n\n", usr.Name, usr.Age, usr.Phone)
}

// Receiver Function
func (usr User) User_Details() {
	fmt.Printf("Name: %s\nAge: %d\nPhone: %s\n\n", usr.Name, usr.Age, usr.Phone)
}

func (usr User) Call(money int) {
	fmt.Printf("%s Won %d$", usr.Name, money)
}

func main() {
	var user1 User
	user1 = User{
		Name:  "Farhan",
		Age:   23,
		Phone: "01412345678",
	}

	Print_User_Details(user1)

	user2 := User{
		Name:  "Nadim",
		Age:   19,
		Phone: "01912345678",
	}

	user2.User_Details()
	user2.Call(300)
}

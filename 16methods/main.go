package main

import "fmt"

// NO INHERITANCE, NO SUPER,PARENT IN GO

// FUNCTIONS INSIDE STRUCTS ARE METHODS
// THATS IT

func main() {
	Akshat := User{"Aksh", "ak@test", true, 22}
	fmt.Println((Akshat))
	Akshat.GetStatus()
	Akshat.GetEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is Active :", u.Status)
}

func (u User) GetEmail() {
	fmt.Println("Email is :", u.Email)
} 

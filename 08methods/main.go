package main

import "fmt"

func main() {
	obj := User{Name: "Bhaskar", Email: "5b5uK@example.com", Active: true, Age: 22}

	obj.GetName()
}

type User struct {
	Name   string
	Email  string
	Active bool
	Age    int
}

func (u User) GetName() {
	fmt.Println("Name is ", u.Name)
}

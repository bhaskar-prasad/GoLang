package main

import "fmt"

func main() {
	Bhaskar := User{
		Name:   "Bhaskar",
		Email:  "5b5uK@example.com",
		Active: true,
		Age:    22,
	}
	fmt.Printf("%+v", Bhaskar)
}

type User struct {
	Name   string
	Email  string
	Active bool
	Age    int
}

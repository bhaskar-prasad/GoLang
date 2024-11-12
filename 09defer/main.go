package main

import "fmt"

func main() {
	defer foo()
	defer fmt.Println("world")
	fmt.Println("Hello")
}

func foo() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

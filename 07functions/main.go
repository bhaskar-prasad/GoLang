package main

import "fmt"

func main() {

	total, message := add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(total, message)
}

func add(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Hello"
}

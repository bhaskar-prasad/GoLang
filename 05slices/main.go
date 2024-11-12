package main

import "fmt"

func main() {
	var numbers = make([]int, 3)

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	fmt.Println("the slice is", numbers)
	numbers = append(numbers[:1], numbers[2:]...)

	fmt.Println("the slice is", numbers)
	var abc = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("the slice is", abc)
}

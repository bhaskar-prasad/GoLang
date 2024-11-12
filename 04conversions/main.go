package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your rating: ")
	input, _ := reader.ReadString('\n')
	newInput, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	fmt.Println("Your Rating is,", newInput+1)

}

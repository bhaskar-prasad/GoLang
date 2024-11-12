package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files")
	content := "This needs to go in a file"
	file, err := os.Create("./abc.txt")
	checkNilError(err)
	defer file.Close()
	length, err := file.WriteString(content)
	checkNilError(err)
	fmt.Println("length is", length)
	data, err := os.ReadFile("./abc.txt")
	checkNilError(err)
	fmt.Println(string(data))

	readFile("./abc.txt")
}

func readFile(filename string) {
	databytes, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println(string(databytes))
}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

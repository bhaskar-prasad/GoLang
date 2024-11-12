package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "http://127.0.0.1:5500/abc.html"

func main() {
	resp, err := http.Get(url)
	checkNilError(err)
	fmt.Println("Status code is", resp.StatusCode)
	defer resp.Body.Close()
	Content, err := io.ReadAll(resp.Body)
	checkNilError(err)
	fileName := "./index.html"
	writeFile(fileName, string(Content))
}
func writeFile(filename string, content string) {
	file, err := os.Create(filename)
	checkNilError(err)
	defer file.Close()
	data, err := file.WriteString(content)
	checkNilError(err)
	fmt.Println("File created with length", data)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

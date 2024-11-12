package main

import (
	"fmt"
	"net/url"
)

const urlName string = "http://127.0.0.1:5500/abc.html/hahaa?course=1&hah=2"

func main() {
	fmt.Println("URL handling")
	res, _ := url.Parse(urlName)
	// fmt.Println(res.Path)
	// fmt.Println(res.RawQuery)
	q := res.Query()
	fmt.Println(q)
	fmt.Println(q["course"])
}

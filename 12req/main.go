package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const Url = "http://localhost:8000/"

func main() {
	fmt.Println("Hello World")
	PerfromGet(Url + "get")
	PerformPost(Url + "post")
	PostForm(Url + "postform")
}

func PerfromGet(url string) {
	// getUrl := url + "get"
	res, err := http.Get(url)
	CheckNilError(err)

	defer res.Body.Close()
	fmt.Println("Status code is", res.StatusCode)
	fmt.Println("Content Length is", res.ContentLength)
	var responseString strings.Builder

	content, err := io.ReadAll(res.Body)
	CheckNilError(err)
	responseString.Write(content)
	fmt.Println(responseString.String())

}

func PerformPost(url string) {
	// postUrl := url + "post"
	requestBody := strings.NewReader((`
	{
		"name":"Bhaskar",
		"email":"5b5uK@example.com",
		"age":22
	}
	`))
	res, err := http.Post(url, "application/json", requestBody)
	CheckNilError(err)
	defer res.Body.Close()
	fmt.Println(res.Body)
	contentLength, err := io.ReadAll(res.Body)
	CheckNilError(err)
	fmt.Println(string(contentLength))
}

func PostForm(postUrl string) {
	// FormUrl := postUrl + "postform"
	data := url.Values{}
	data.Add("name", "Bhaskar")
	data.Add("age", "22")

	res, err := http.PostForm(postUrl, data)
	CheckNilError(err)
	defer res.Body.Close()
	contentLength, err := io.ReadAll(res.Body)
	CheckNilError(err)
	fmt.Println(string(contentLength))
}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}

}

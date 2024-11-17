package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Grade    int      `json:"grade"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Hello, World!")
	JsonEncode()
}

func JsonEncode() {
	jsonCourse := []course{
		{
			Name:     "Go",
			Grade:    90,
			Password: "123456",
			Tags:     []string{"Go", "Golang"},
		},
		{
			Name:     "Java",
			Grade:    80,
			Password: "123456",
			Tags:     []string{"Java", "JavaEE"},
		},
	}
	// jsonData, err := json.Marshal(jsonCourse)
	jsonData, err := json.MarshalIndent(jsonCourse, "", "\t")
	CheckNilError(err)
	fmt.Printf("%s\n", jsonData)
}
func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}

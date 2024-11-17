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
	// JsonEncode()
	DecodeJson()

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

func DecodeJson() {
	jsonData := []byte(`
		{
		"coursename": "Go",
		"grade": 90,
		"tags": ["Go", "Golang"]
		}
	`)
	var jsonCourse course
	checkValid := json.Valid(jsonData)
	if checkValid {
		json.Unmarshal(jsonData, &jsonCourse)
		fmt.Printf("%#v\n", jsonCourse)
	} else {
		fmt.Println("Invalid JSON")
	}

}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}

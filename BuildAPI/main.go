package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/bhaskar-prasad/buildapi/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("API")
	r := mux.NewRouter()
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/courses", GetCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")
	r.HandleFunc("/course", CreateOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")
	defer log.Fatal(http.ListenAndServe(":8000", r))
}

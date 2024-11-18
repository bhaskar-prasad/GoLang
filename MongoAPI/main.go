package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhaskar-prasad/mongoapi/router"
)

func main() {
	r := router.Router()
	fmt.Println("Server Starting")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server Started")
}

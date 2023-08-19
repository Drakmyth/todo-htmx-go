package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")

	registerTaskEndpoints()
	http.Handle("/", http.FileServer(http.Dir("public")))

	log.Fatal(http.ListenAndServe(":8000", nil))
}

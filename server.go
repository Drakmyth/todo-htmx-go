package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")

	h1 := http.FileServer(http.Dir("public"))

	http.Handle("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

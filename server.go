package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")

	h1 := http.FileServer(http.Dir("public"))
	h2 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<p>I've Been Clicked!</p>"))
	}

	http.HandleFunc("/clicked", h2)
	http.Handle("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

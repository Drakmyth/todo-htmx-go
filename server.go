package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")

	addHandler := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		item := r.Form.Get("item")
		fmt.Print(item)
		fmt.Fprintf(w, "<li>%s</li>", item)
	}
	rootHandler := http.FileServer(http.Dir("public"))

	http.HandleFunc("/add", addHandler)
	http.Handle("/", rootHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

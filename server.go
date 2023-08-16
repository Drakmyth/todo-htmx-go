package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	fmt.Println("Hello World!")

	addHandler := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		item := r.Form.Get("item")
		tmpl, _ := template.ParseFiles("./templates/item.tmpl.html")
		tmpl.Execute(w, item)
	}
	rootHandler := http.FileServer(http.Dir("public"))

	http.HandleFunc("/add", addHandler)
	http.Handle("/", rootHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

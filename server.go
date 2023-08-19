package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID
	Description string
}

func main() {
	fmt.Println("Hello World!")

	addHandler := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		description := r.Form.Get("description")
		task := Task{
			Id:          uuid.New(),
			Description: description,
		}
		tmpl, _ := template.ParseFiles("./templates/item.tmpl.html")
		tmpl.Execute(w, task)
	}

	http.HandleFunc("/add", addHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))

	log.Fatal(http.ListenAndServe(":8000", nil))
}

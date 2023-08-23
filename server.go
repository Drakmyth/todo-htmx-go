package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/jba/muxpatterns"
)

var taskMux *muxpatterns.ServeMux = nil
var tmpl *template.Template = nil

func main() {
	fmt.Println("Hello World!")

	mux := muxpatterns.NewServeMux()
	taskMux = NewTaskMux()

	tmpl, _ = template.ParseGlob("./templates/*.tmpl.html")

	mux.Handle("/tasks/", http.StripPrefix("/tasks", taskMux))
	mux.Handle("/", http.FileServer(http.Dir("public")))

	log.Fatal(http.ListenAndServe(":8000", mux))
}

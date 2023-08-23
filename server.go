package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jba/muxpatterns"
)

var taskMux *muxpatterns.ServeMux = nil

func main() {
	fmt.Println("Hello World!")

	mux := muxpatterns.NewServeMux()
	taskMux = NewTaskMux()

	mux.Handle("/tasks/", http.StripPrefix("/tasks", taskMux))
	mux.Handle("/", http.FileServer(http.Dir("public")))

	log.Fatal(http.ListenAndServe(":8000", mux))
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")

	mux := http.NewServeMux()
	taskMux := NewTaskMux()

	mux.Handle("/tasks/", http.StripPrefix("/tasks", taskMux))
	mux.Handle("/", http.FileServer(http.Dir("public")))

	log.Fatal(http.ListenAndServe(":8000", mux))
}

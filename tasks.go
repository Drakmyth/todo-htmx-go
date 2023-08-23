package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID
	Description string
}

var tasks = make(map[uuid.UUID]Task)

func NewTaskMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", taskRootHandler)
	return mux
}

func createTask(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Hello!")
	r.ParseForm()

	id := uuid.New()
	description := r.Form.Get("description")
	task := Task{
		Id:          id,
		Description: description,
	}
	tasks[id] = task

	tmpl, _ := template.ParseFiles("./templates/item.tmpl.html")
	tmpl.Execute(w, task)
}

func taskRootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createTask(w, r)
	default:
		methodNotAllowedHandler(w, r)
	}
}

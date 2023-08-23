package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/google/uuid"
	"golang.org/x/exp/maps"
)

type Task struct {
	Id          uuid.UUID
	Description string
}

type TaskList struct {
	Tasks []Task
}

var tasks = make(map[uuid.UUID]Task)

func NewTaskMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", taskRootHandler)
	return mux
}

func createTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id := uuid.New()
	description := r.Form.Get("description")
	task := Task{
		Id:          id,
		Description: description,
	}
	tasks[id] = task

	fmt.Printf("Created %s\n", task)

	tmpl, _ := template.ParseFiles("./templates/task.tmpl.html")
	tmpl.ExecuteTemplate(w, "task", task)
}

func getTaskList(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Retrieving %d tasks...\n", len(tasks))

	taskList := TaskList{
		Tasks: maps.Values(tasks),
	}

	tmpl, _ := template.ParseFiles("./templates/task.tmpl.html")
	tmpl.ExecuteTemplate(w, "task-list", taskList)
}

func taskRootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createTask(w, r)
	case http.MethodGet:
		getTaskList(w, r)
	default:
		methodNotAllowedHandler(w, r)
	}
}

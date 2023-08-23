package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/google/uuid"
	"github.com/jba/muxpatterns"
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

func NewTaskMux() *muxpatterns.ServeMux {
	mux := muxpatterns.NewServeMux()

	mux.HandleFunc("DELETE /{taskId}", deleteTask)
	mux.HandleFunc("POST /", createTask)
	mux.HandleFunc("GET /", getTaskList)
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

func deleteTask(w http.ResponseWriter, r *http.Request) {
	id, _ := uuid.Parse(taskMux.PathValue(r, "taskId"))

	fmt.Printf("Deleting %s\n", id)
	delete(tasks, id)

}

package main

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/jba/muxpatterns"
	"golang.org/x/exp/maps"
)

type Task struct {
	Id          uuid.UUID
	Created     time.Time
	Description string
}

type TaskList struct {
	Tasks []Task
}

var tasks = make(map[uuid.UUID]Task)

func NewTaskMux() *muxpatterns.ServeMux {
	mux := muxpatterns.NewServeMux()

	mux.HandleFunc("GET /{taskId}", getTask)
	mux.HandleFunc("DELETE /{taskId}", deleteTask)
	mux.HandleFunc("GET /{taskId}/edit", getEditTask)
	mux.HandleFunc("POST /", createTask)
	mux.HandleFunc("GET /{$}", getTaskList)
	return mux
}

func createTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id := uuid.New()
	now := time.Now()
	description := r.Form.Get("description")
	task := Task{
		Id:          id,
		Created:     now,
		Description: description,
	}
	tasks[id] = task

	fmt.Printf("Created %s\n", task)

	tmpl.ExecuteTemplate(w, "task", task)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	taskId, _ := uuid.Parse(taskMux.PathValue(r, "taskId"))

	fmt.Printf("Retrieving %s\n", taskId)
	task := tasks[taskId]

	tmpl.ExecuteTemplate(w, "task", task)
}

func getEditTask(w http.ResponseWriter, r *http.Request) {
	taskId, _ := uuid.Parse(taskMux.PathValue(r, "taskId"))

	fmt.Printf("Editing %s\n", taskId)
	task := tasks[taskId]

	tmpl.ExecuteTemplate(w, "edit-task", task)
}

func getTaskList(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Retrieving %d tasks...\n", len(tasks))

	taskValues := maps.Values(tasks)
	slices.SortFunc(taskValues, func(a, b Task) int { return a.Created.Compare(b.Created) })

	taskList := TaskList{
		Tasks: taskValues,
	}

	tmpl.ExecuteTemplate(w, "task-list", taskList)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	taskId, _ := uuid.Parse(taskMux.PathValue(r, "taskId"))

	fmt.Printf("Deleting %s\n", taskId)
	delete(tasks, taskId)

}

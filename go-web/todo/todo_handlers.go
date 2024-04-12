package todo

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type appState struct {
	Todos  []ToDo
	nextId int
}

const (
	TodoTmpl    = "todo"
	TodoRowTmpl = "todoRow"
)

var states = appState{
	Todos:  make([]ToDo, 0),
	nextId: 0,
}

var templates = template.Must(template.ParseGlob("static/templates/*.gohtml"))

func todoPageHandler(w http.ResponseWriter, r *http.Request) {
	render(w, TodoTmpl, states)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	task := r.FormValue("task")
	description := r.FormValue("description")

	states.Todos = append(states.Todos, ToDo{getNextId(),
		task, description, false})

	render(w, TodoRowTmpl, states)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for i, t := range states.Todos {
		if t.Id == id {
			states.Todos = append(states.Todos[:i], states.Todos[i+1:]...)
			break
		}
	}

	render(w, TodoRowTmpl, states)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for i, t := range states.Todos {
		if t.Id == id {
			states.Todos[i].Completed = true
		}
	}

	render(w, TodoRowTmpl, states)
}

func render(w http.ResponseWriter, name string, data appState) {
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

func getNextId() string {
	states.nextId++
	return fmt.Sprintf("%d", states.nextId)
}

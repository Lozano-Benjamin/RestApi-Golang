package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Aca creamos un struct para las tareas
type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

//Aca todas las tareas
type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a mi API xd")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	//Con esto explicamos al servidor que tipo de Dato enviamos
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTasks(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Insert a Valid Task")
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func main() {

	//Creamos aca un enrutador, donde si o si la ruta tiene que estar bien escrita
	router := mux.NewRouter().StrictSlash(true)

	//esta funcion sirve para indicar segun la ruta, que funcion cumplir
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTasks).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}

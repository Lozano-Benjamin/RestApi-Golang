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

//Aca todas las tareas y la definimos como lista
type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
}

//Funcion del index
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a mi API xd")
}

//Funcion para pedir las tareas
func getTasks(w http.ResponseWriter, r *http.Request) {
	//Con esto explicamos al servidor que tipo de Dato enviamos
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

//Funcion para crear la tarea
func createTasks(w http.ResponseWriter, r *http.Request) {
	//Creamos variable tarea
	var newTask task
	//Aca creamos una variable que lee el request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//En caso de error, te sale un aviso
		fmt.Fprint(w, "Insert a Valid Task")
	}

	json.Unmarshal(reqBody, &newTask)

	//Aca hacemos un ID automatico, (Cuando añada la base de datos lo cambiare)
	newTask.ID = len(tasks) + 1
	//agregamos la tarea a la lista
	tasks = append(tasks, newTask)

	//Mandamos el tipo de contenido, y avisos al servidor.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func main() {

	//Creamos aca un enrutador, donde si o si la ruta tiene que estar bien escrita
	router := mux.NewRouter().StrictSlash(true)

	//Creamos rutas con sus metodos
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTasks).Methods("POST")

	//Definimos el puerto
	log.Fatal(http.ListenAndServe(":3000", router))
}

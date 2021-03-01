package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

//Funcion para pedir una tarea individual
func getTask(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	taskID, err :=strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	//Se busca entre las tasks el ID solicitado y luego se muestra en forma de JSON
	for _, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

//funcion para eliminar una tarea
func deleteTask(w http.ResponseWriter, r *http.Request){
	//definimos variable de vars que devuelve las variables de ruta
	vars := mux.Vars(r)

	taskID, err := strconv.Atoi(vars["id"])
	if err != nil{
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	//Se elimina la task a la lista, guardando todas las que estan hasta su indice, y la que le sigue en adelante.
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i + 1:] ...)
			fmt.Fprintf(w, "The task with ID %v has been removed succesfully", taskID)
		}
	}
}

//Funcion para actualizar una tarea
func updateTask(w http.ResponseWriter, r *http.Request){
	//definimos variable de vars que devuelve las variables de ruta
	vars := mux.Vars(r)
	//convertimos la variable del id a ints
	taskID, err := strconv.Atoi(vars["id"])

	//Creamos una variable donde almacenaremos la nueva tarea 
	var updatedTask task


	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	//Creamos una funcion que lee todo el body del request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w,"Please enter Valid Data")
	}

	//Desarma el Json y lo hace una struct
	json.Unmarshal(reqBody, &updatedTask)

	
	//Se busca entre todas las tasks una task con el ID solicitado
	for i, task := range tasks {
		if task.ID == taskID {
			//Se elimina la task a la lista, guardando todas las que estan hasta su indice, y la que le sigue en adelante.
			tasks = append(tasks[:i], tasks[i + 1:]...)

			//El id se mantiene
			updatedTask.ID = taskID
			//Se agrega nueva task
			tasks = append(tasks, updatedTask)

			//Aviso de que la task se cambio con exito
			fmt.Fprintf(w, "The task with ID %v has been succesfully updated", taskID)
		}
	}
}

func main() {

	//Creamos aca un enrutador, donde si o si la ruta tiene que estar bien escrita
	router := mux.NewRouter().StrictSlash(true)

	//Creamos rutas con sus metodos
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTasks).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask ).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	//Definimos el puerto
	log.Fatal(http.ListenAndServe(":3000", router))
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/task/{id}/", handleTaskByID)
	mux.HandleFunc("/task/{id}/status", handleTaskStatusByID)
	mux.HandleFunc("GET /home", handleHome)

	log.Default().Println("Booting server")
	http.ListenAndServe("localhost:8090", mux)
}

// Handlers
func handleTaskByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Printf("200 GET /task/%s", id)
	fmt.Fprintf(w, "handling task with id=%v\n", id)
}

func handleTaskStatusByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Printf("200 GET /task/%s/status", id)
	fmt.Fprintf(w, "handling task status with id=%v\n", id)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	log.Println("200 GET /home")
	fmt.Fprint(w, "home\n")
}

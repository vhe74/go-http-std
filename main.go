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

	loggedMux := withLog(mux)

	log.Default().Println("Booting server")
	http.ListenAndServe("localhost:8090", loggedMux)
}

// Middleware
func withLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// Handlers
func handleTaskByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "handling task with id=%v\n", id)
}

func handleTaskStatusByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "handling task status with id=%v\n", id)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home\n")
}

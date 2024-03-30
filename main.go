package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/task/{id}/", handleTaskByID)
	mux.HandleFunc("/task/{id}/status", handleTaskStatusByID)
	mux.HandleFunc("GET /home", handleHome)
	mux.HandleFunc("GET /wait/{waitsecs}", handleWait)

	loggedMux := withLog(mux)

	log.Default().Println("Booting server")
	http.ListenAndServe("localhost:8090", loggedMux)
}

// Middleware
func withLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, time.Since(startTime))
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

func handleWait(w http.ResponseWriter, r *http.Request) {
	waitsecs, err := strconv.Atoi(r.PathValue("waitsecs"))
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error getting wait value \n")
		return
	}
	time.Sleep(time.Duration(waitsecs) * time.Second)
	fmt.Fprintf(w, "Waitted %d seconds\n", waitsecs)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {

	httpServer := NewHttpServer("localhost:8090")

	httpServer.AddHandleFunc("/task/{id}/", handleTaskByID)
	httpServer.AddHandleFunc("/task/{id}/status", handleTaskStatusByID)
	httpServer.AddHandleFunc("GET /{$}", handleHome)
	httpServer.AddHandleFunc("GET /wait/{waitsecs}", handleWait)
	httpServer.AddHandleFunc("GET /static/{filename...}", handleServeFile)

	httpServer.Run()

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

func handleServeFile(w http.ResponseWriter, r *http.Request) {
	path := "static/" + r.PathValue("filename")
	http.ServeFile(w, r, path)
}

type HttpServer struct {
	mux     *http.ServeMux
	address string
}

func NewHttpServer(address string) HttpServer {
	return HttpServer{
		mux:     http.NewServeMux(),
		address: address,
	}
}

func (h *HttpServer) AddHandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	h.mux.HandleFunc(pattern, handler)
}

func (h *HttpServer) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		//log.Printf("%+v", r)
		//log.Printf("%+v", r.Context())
		//log.Printf("%+v", w.Header())
		log.Printf("[%s] %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, time.Since(startTime))
	})
}

func (h *HttpServer) Run() {
	log.Println("Booting server")
	loggedMux := h.Logger(h.mux)
	http.ListenAndServe(h.address, loggedMux)
}

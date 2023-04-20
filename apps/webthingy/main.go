package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Page not found")
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		pageNotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello from some kinda web thingy")
}

type WebThingy struct {
	server *http.Server
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v %v %v", r.Method, r.URL, r.Proto, r.Response)
		next.ServeHTTP(w, r)
	})
}

func NewWebThingy() *WebThingy {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)

	handler := logging(mux)
	return &WebThingy{
		server: &http.Server{
			Addr:           ":8080",
			Handler:        handler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (w *WebThingy) Run() {
	log.Printf("Listening on HTTP port '%v'\n", w.server.Addr)
	log.Fatal(w.server.ListenAndServe())
}

func main() {
	NewWebThingy().Run()
}

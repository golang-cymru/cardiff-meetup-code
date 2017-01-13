package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(w, "Hi there, I love surveys!")
}

func survey(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	reference := vars["reference"]
	fmt.Fprintf(w, "Hi there, I love %s!", reference)
}

func main() {

	// Configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Print("$PORT defaulted to " + port)
	}

	// Mux
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/{reference}", survey)
	http.Handle("/", r)

	// Run
	http.ListenAndServe(":"+port, nil)
}

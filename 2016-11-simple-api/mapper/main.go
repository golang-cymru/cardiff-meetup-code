package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

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

func wool(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	material := vars["material"]
	count := vars["count"]

	type Inventory struct {
		Material string
		Count    string
	}
	sweaters := Inventory{material, count}
	tmpl, err := template.ParseFiles("inventory.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, sweaters)
	if err != nil {
		panic(err)
	}
}

func mapper(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	lat := vars["lat"]
	long := vars["long"]

	type Location struct {
		Lat  float64
		Long float64
	}

	latF, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		panic(err)
	}
	longF, err := strconv.ParseFloat(long, 64)
	if err != nil {
		panic(err)
	}
	loc := Location{latF, longF}
	tmpl, err := template.ParseFiles("map.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, loc)
	if err != nil {
		panic(err)
	}
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
	r.HandleFunc("/inventory/{material}/{count}", wool)
	r.HandleFunc("/map/{lat}/{long}", mapper)
	http.Handle("/", r)

	// Run
	http.ListenAndServe(":"+port, nil)
}

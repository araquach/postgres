package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	tplHome *template.Template
	tplCreate *template.Template
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplHome.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	tplHome = template.Must(template.ParseFiles("templates/index.gohtml"))

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")

	http.ListenAndServe(":" + port, r)
}
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"html/template"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	dbport     = 5432
	user     = "adam-macbook"
	password = "Blonde123"
	dbname   = "postgres"
)

var (
	tplHome *template.Template
)

type Applicant struct{
	gorm.Model
	Name string
	Mobile string
	Position string
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplHome.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	// port := os.Getenv("PORT")
	   port := "5050"
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	tplHome = template.Must(template.ParseFiles("templates/index.gohtml"))

	r := mux.NewRouter()
	r.HandleFunc("/", home)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, dbport, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	db.AutoMigrate(&Applicant{})

	http.Handle("/", r)
	http.ListenAndServe(":" + port, r)
}
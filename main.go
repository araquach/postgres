package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	tplHome *template.Template
	tplCreate *template.Template
)

type Applicant struct{
	gorm.Model
	Name string
	Mobile string
	Position string
}

func home(w http.ResponseWriter, r *http.Request) {
	applicant := Applicant{
		Name: "Izzy Lamb",
		Mobile: "07921806884",
		Position: "Senior Stylist",
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tplHome.Execute(w, applicant); err != nil {
		panic(err)
	}
}

func apply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplCreate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var (
		host     = os.Getenv("DB_HOST")
		dbport   = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	tplHome = template.Must(template.ParseFiles("templates/index.gohtml"))
	tplCreate = template.Must(template.ParseFiles("templates/create.gohtml"))

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/apply", apply).Methods("GET")
	//r.HandleFunc("/apply", apply).Methods("POST")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, dbport, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	db.AutoMigrate(&Applicant{})

	http.ListenAndServe(":" + port, r)
}
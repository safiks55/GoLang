package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB

func main() {

	db, err := sql.Open("mysql", "SeD:1379@tcp(127.0.0.1:3306)/stud")
	if err != nil {
		fmt.Println("error connecting to the db")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection to the db")
		panic(err.Error())
	}

	r := mux.NewRouter()
	r.HandleFunc("/allStudents", getAllStudents).Methods("GET")
	r.HandleFunc("/student/{id}", getStudent).Methods("GET")
	r.HandleFunc("/student", addStudent).Methods("POST")
	r.HandleFunc("/student/{id}", updateStudent).Methods("PUT")
	r.HandleFunc("/student/{id}", deleteStudent).Methods("DELETE")

	fmt.Printf("Starting server at PORT:8888\n")
	log.Fatal(http.ListenAndServe(":8888", r))
}

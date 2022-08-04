package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var students = []student{
	{firstName: "sam", lastName: "smith", id: "1ss", email: "sam1@ymail.com"},
	{firstName: "jason", lastName: "darulo", id: "jd2", email: "jd2@ymail.com"},
}

// get all students in db
func getAllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllStudents")
	json.NewEncoder(w).Encode(students)
}

// we retrieve the students from the route using mux.Vars(). Then iterate through the slice and return only the requested student.
func getStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchId := vars["id"]
	for _, student := range students {
		if student.id == searchId {
			json.NewEncoder(w).Encode(student)
		}
	}
}

// we recieve the post request then we parse json data using unmarshalling and then append it to our list of students.
func addStudent(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var temp student
	json.Unmarshal(reqBody, &temp)
	students = append(students, temp)

	json.NewEncoder(w).Encode(students)
}

//will update a student if the id matches one of the students in the slice. It will take the PUT request, decode it and store it in update variable
func updateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchId := vars["id"]

	for i, student := range students {
		if student.id == searchId {
			students = append(students[:i], students[i+1:]...)
		}
	}

	var updateStud student
	json.NewDecoder(r.Body).Decode(&updateStud)
	students = append(students, updateStud)
	fmt.Println("Endpoint hit: UpdateStudent")
	json.NewEncoder(w).Encode(updateStud)
	return
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	searchID := vars["id"]

	for i, student := range students {
		if student.id == searchID {
			students = append(students[:i], students[i+1:]...)
		}
	}
}

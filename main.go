package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type student struct {
	ID string `json: "ID"`
	Name string `json: "Name"`
	Age int `json: "Age"`
}

type allStudents []student

var students = allStudents{
	{
		ID: "1",
		Name: "Suande",
		Age: 22,
	},
}

func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/students", getAllStudents).Methods("GET")
	router.HandleFunc("/students/{id}", getStudentsById).Methods("GET")
	router.HandleFunc("/student", createStudent).Methods("POST")
	router.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")
	router.HandleFunc("/students/{id}", updateStudent).Methods("PUT")

	log.Fatal(http.ListenAndServe(":9000", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to students api")
}

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "oke");
	json.NewEncoder(w).Encode(students)
}

func getStudentsById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for _,student := range students {
		if student.ID == id {
			json.NewEncoder(w).Encode(student)
		}
	}
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	var newStudent student
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "something wrong")
	}

	json.Unmarshal(reqBody, &newStudent)
	students = append(students, newStudent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newStudent)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			fmt.Fprintf(w, "the studetn with id %v has been deleted succefully", id)
		}
	}
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updateStudent student

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "some thing wrong")
	}

	json.Unmarshal(reqBody, &updateStudent)

	for i, student := range students {
		if student.ID == id {
			student.Name = updateStudent.Name
			student.Age = updateStudent.Age

			students = append(students[:i], student)
			json.NewEncoder(w).Encode(student)
		}
	}
}



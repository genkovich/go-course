package main

import (
	"course/hw9"
	"course/hw9/task"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("Server start")
	r := mux.NewRouter()

	taskHandler := task.NewHandler()
	studentHandler := hw9.NewHandler()

	r.HandleFunc("/tasks", taskHandler.GetTasksByDate)
	//r.HandleFunc("/class/{class:[A-Z]+}", studentHandler.GetStudentsByClass)
	//r.HandleFunc("/class/{class:[A-Z]+}/statistic", studentHandler.GetClassStatistic)
	//r.HandleFunc("/student/{id:[0-9]+}", studentHandler.GetStudentById)

	r.Handle("/class/{class:[A-Z]+}", studentHandler.AuthMiddleware(http.HandlerFunc(studentHandler.GetStudentsByClass)))
	r.Handle("/class/{class:[A-Z]+}/statistic", studentHandler.AuthMiddleware(http.HandlerFunc(studentHandler.GetClassStatistic)))
	r.Handle("/student/{id:[0-9]+}", studentHandler.AuthMiddleware(http.HandlerFunc(studentHandler.GetStudentById)))

	http.ListenAndServe(":8082", r)
}

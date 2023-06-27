package main

import (
	"course/hw9/task"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("Server start")
	r := mux.NewRouter()

	taskHandler := task.NewHandler()

	r.HandleFunc("/tasks", taskHandler.GetTasksByDate)
	http.ListenAndServe(":8082", r)
}

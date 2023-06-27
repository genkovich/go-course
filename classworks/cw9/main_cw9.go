package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Написати вебсервер, який по GET отримує 2 числа і знак (*/+-), а відповідає результатом цієї операції

func main() {

	fmt.Println("Server start")
	r := mux.NewRouter()

	r.HandleFunc("/count", mathCount)
	http.ListenAndServe(":8082", r)
}

type RequestBody struct {
	Number1 int    `json:"number1"`
	Number2 int    `json:"number2"`
	Sign    string `json:"sign"`
}

func mathCount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Request")

	reqBody := RequestBody{}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var result int

	switch reqBody.Sign {
	case "+":
		result = reqBody.Number1 + reqBody.Number2
	case "-":
		result = reqBody.Number1 - reqBody.Number2
	case "*":
		result = reqBody.Number1 * reqBody.Number2
	case "/":
		result = reqBody.Number1 / reqBody.Number2
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

	jsonMessage, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonMessage)
	if err != nil {
		return
	}
}

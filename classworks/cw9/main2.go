package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	fmt.Println("Server start")
	r := mux.NewRouter()

	s := &Storage{}

	h := &CityHandler{s}

	r.Handle("/cities", authMiddleware(http.HandlerFunc(h.getCities)))
	r.Handle("/new-city", authMiddleware(http.HandlerFunc(h.postCities)))

	http.ListenAndServe(":8083", r)
}

type User struct {
	Username string
	Password string
}

var adminUser = User{
	Username: "admin",
	Password: "admin",
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if username != adminUser.Username || password != adminUser.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *CityHandler) getCities(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Request")

	jsonMessage, err := json.Marshal(h.s.GetCities())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonMessage)
}

type CityHandler struct {
	s *Storage
}

type singleCityReqBody struct {
	City string
}

func (h *CityHandler) postCities(w http.ResponseWriter, r *http.Request) {
	reqBody := &singleCityReqBody{}

	fmt.Println("Handling Request")
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		fmt.Println(err)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.s.CreateCity(reqBody.City)
}

package main

import (
	"course/hw10/translator"
	"course/hw10/weather"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	r := mux.NewRouter()

	weatherRes := &weather.Resource{}
	transRes := &translator.Resource{}

	r.HandleFunc("/weather", weatherRes.GetCityWeather).Methods(http.MethodGet)
	r.HandleFunc("/translate", transRes.Translate).Methods(http.MethodPost)
	log.Info().Msg("Server starting..")
	http.ListenAndServe(":8082", r)
}

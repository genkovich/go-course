package weather

import (
	"course/hw10/helper"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"strings"
)

type Resource struct{}

const apiUrl = "http://api.weatherapi.com/v1/current.json?key={api_key}&q={city}&aqi=no"

func (wr *Resource) GetCityWeather(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("handling weather request")

	apiKey := os.Getenv("WEATHER_API_KEY")

	url := strings.Replace(apiUrl, "{api_key}", apiKey, 1)

	city := r.URL.Query().Get("city")

	if city == "" {
		log.Error().Msg("city is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url = strings.Replace(url, "{city}", city, 1)

	resp, err := http.Get(url)
	if err != nil {
		log.Error().Msg("Failed to get weather")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Error().Msg("Non OK status from weather")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type respBody struct {
		Location map[string]any `json:"location"`
		Current  map[string]any `json:"current"`
	}

	var weather respBody

	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		log.Error().Msgf("cant decode weather %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	helper.RespondWithJson(w, weather)
}

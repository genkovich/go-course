package helper

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, body any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Error().Msg("Something wrong while encode answer")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

package translator

import (
	"bytes"
	"course/hw10/helper"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

type Resource struct{}

const apiUrl = "https://api.lecto.ai/v1/translate/text"

type postBody struct {
	To   string `json:"to"`
	From string `json:"from"`
	Text string `json:"text"`
}

func (t *Resource) Translate(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("handling translate request")

	apiKey := os.Getenv("TRANSLATE_API_KEY")

	var vars postBody
	err := json.NewDecoder(r.Body).Decode(&vars)
	if err != nil {
		log.Error().Msg("Failed to decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bodyData := map[string]interface{}{
		"texts": []string{vars.Text},
		"to":    []string{vars.To},
		"from":  vars.From,
	}

	body, err := json.Marshal(bodyData)
	if err != nil {
		log.Error().Err(err).Msg("Failed to encode request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to translate")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		log.Error().Msg("Non OK status from translate")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type TranslationItem struct {
		To         string   `json:"to"`
		Translated []string `json:"translated"`
	}

	type TranslationResponse struct {
		Translations         []TranslationItem `json:"translations"`
		From                 string            `json:"from"`
		TranslatedCharacters int               `json:"translated_characters"`
	}

	var translations TranslationResponse

	err = json.NewDecoder(resp.Body).Decode(&translations)
	if err != nil {
		log.Error().Msgf("cant decode translations %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	helper.RespondWithJson(w, translations)
}

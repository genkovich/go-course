package hw11

import (
	"bufio"
	"github.com/rs/zerolog/log"
	"os"
	"regexp"
)

func SearchPhoneNumbers() {

	file, err := os.Open("./hw11/sources/numbers.txt")
	if err != nil {
		log.Error().Err(err).Msg("Error open file")
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var phoneNumbers []string

	for fileScanner.Scan() {
		//matched, err := regexp.MatchString(`^[\d()\s.-]{10,}$`, fileScanner.Text())
		matched, err := regexp.MatchString(`^\(?\d{3}[)? |\-.]{0,2}\d{3}[\-?|. ]?\d{4}$`, fileScanner.Text())
		if err != nil {
			log.Error().Err(err).Msg("Error regexp")
			return
		}
		if matched {
			phoneNumbers = append(phoneNumbers, fileScanner.Text())
		}
	}

	log.Info().Msg("Phone numbers:")
	for _, phoneNumber := range phoneNumbers {
		log.Info().Msg(phoneNumber)
	}
}

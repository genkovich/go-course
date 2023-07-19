package hw11

import (
	"bufio"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"regexp"
	"strings"
)

func SearchWords() {
	file, err := os.Open("./hw11/sources/text.txt")
	if err != nil {
		log.Error().Err(err).Msg("Error open file")
		return
	}

	scanner := bufio.NewScanner(file)

	words := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, strings.Fields(line)...)
	}

	regexps := []*regexp.Regexp{
		regexp.MustCompile(`^\p{L}{3,5}$`),
		regexp.MustCompile(`^[аеєиіїоуюяАЕЄИІЇОУЮЯ]\p{L}*[аеєиіїоуюяАЕЄИІЇОУЮЯ]$`),
		regexp.MustCompile(`^\p{Lu}\p{Ll}*$`),
		regexp.MustCompile(`^\p{L}{8,}$`),
	}

	for i, regex := range regexps {
		for _, word := range words {
			if regex.MatchString(word) {
				fmt.Printf("Regex %d matched word: %s\n", i+1, word)
			}
		}
	}
}

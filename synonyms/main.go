package main

import (
	"bufio"
	"fmt"
	"github.com/daiLlew/go-cmd-line-tool/thesaurus"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("BHT_API_KEY")
	if len(apiKey) == 0 {
		log.Fatalln("Could not find BHT_API_KEY in env. Exiting.")
		os.Exit(0)
	}

	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()

		if word == "" {
			continue
		}

		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalln("Failed when looking for synonyms for "+word, err)
		}
		if len(syns) == 0 {
			log.Fatalln("Couldn't find any synonyms for " + word)
		}
		for _, s := range syns {
			fmt.Println(s)
		}
	}
}

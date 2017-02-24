package thesaurus

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const BigHugeURLFMT = "http://words.bighugelabs.com/api/2/%s/%s/json"

type BigHuge struct {
	APIKey string
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	url := fmt.Sprintf(BigHugeURLFMT, b.APIKey, term)
	response, err := http.Get(url)
	if err != nil {
		return syns, errors.New("bigHuge: Failed when looking for sysnonyms for " + term + ": " + err.Error())
	}
	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}
	if data.Noun != nil {
		syns = append(syns, data.Noun.Syn...)
	}
	if data.Verb != nil {
		syns = append(syns, data.Verb.Syn...)
	}
	return syns, nil
}

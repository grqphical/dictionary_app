package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Phonetic struct {
	Text string `json:"text"`
}

type Definition struct {
	Definition string `json:"definition"`
	Example    string `jsonj:"example"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
}

type DictionaryResult struct {
	Word      string     `json:"word"`
	Phonetics []Phonetic `json:"phonetics"`
	Meanings  []Meaning  `json:"meanings"`
}

// / Get the definition from the dictionary API and convert it into a GO struct
func GetDefinition(word string) (DictionaryResult, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word))

	if err != nil {
		return DictionaryResult{}, err
	}

	if resp.StatusCode == 404 {
		return DictionaryResult{}, fmt.Errorf("Could not find definition for word")
	}

	output := make([]DictionaryResult, 0)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return DictionaryResult{}, fmt.Errorf("Failed to parse request")
	}

	err = json.Unmarshal(data, &output)
	if err != nil {
		return DictionaryResult{}, err
	}

	result := output[0]

	return result, nil
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Cache map[string]*DictionaryResult

func NewCache() Cache {
	return Cache{}
}

func CacheFromFile(file string) (Cache, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return NewCache(), fmt.Errorf("Cannot read file")
	}

	cache := NewCache()

	err = json.Unmarshal(data, &cache)
	if err != nil {
		return cache, fmt.Errorf("Could not deserialize data")
	}

	return cache, nil
}

func (c Cache) AddItem(word string, item *DictionaryResult) {
	// Check if word is already cached and if so return as we
	// do not need to add it again
	_, exists := c[word]
	if exists {
		return
	}

	c[word] = item
}

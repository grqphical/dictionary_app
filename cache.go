package main

type Cache map[string]*DictionaryResult

func NewCache() Cache {
	return Cache{}
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

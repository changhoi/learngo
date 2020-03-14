package mydict

import "errors"

// Dictionary Type
type Dictionary map[string]string

var errNotFound = errors.New("Not founded")
var errWordExists = errors.New("Word exists")
var errCantUpdate = errors.New("Can't Update non existing word")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if exist {
		return value, nil
	}

	return "", errNotFound
}

// Add one word
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err != nil {
		d[word] = def
		return nil
	}
	return errWordExists
}

// Update word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

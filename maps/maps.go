package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("word not in dictionary")
var ErrWordExists = errors.New("word already in dictionary")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	// A map value is a pointer to a runtime.hmap structure
	// (no need for pointer operator!)
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

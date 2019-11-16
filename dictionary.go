package main

type Dictionary map[string]string
type DictionaryError string

var (
	ErrorNotFound   = DictionaryError("could not find the word you were looking for")
	ErrorWordExists = DictionaryError("cannot add word because it already exists")
)

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}

	return nil
}

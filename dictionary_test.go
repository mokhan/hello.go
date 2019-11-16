package main

import (
	"testing"
)

// https://github.com/quii/learn-go-with-tests/blob/master/maps.md
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"
		err := dictionary.Add(key, value)

		assertNil(t, err)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "new test")

		assertError(t, err, ErrorWordExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)

		assertNil(t, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)

		assertError(t, err, ErrorWordDoesNotExist)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, key, expected string) {
	t.Helper()

	actual, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added key:", err)
	}

	if expected != actual {
		t.Errorf("actual %q want %q", actual, expected)
	}
}

func assertNil(t *testing.T, actual error) {
	if actual != nil {
		t.Errorf("expected nil got %q", actual)
	}
}

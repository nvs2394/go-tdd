package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is test"}

	t.Run("know word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is test"

		assertStrings(t, got, want)
	})

	t.Run("unknow word", func(t *testing.T) {
		_, err := dictionary.Search("unknow")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new word")

		assertError(t, err, ErrorWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
		assertError(t, err, nil)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

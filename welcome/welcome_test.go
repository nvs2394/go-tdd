package main

import "testing"

func TestWelcome(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying welcome to people", func(t *testing.T) {
		got := Welcome("Son", "")
		want := "Hello, Son"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying welcome to people in other language", func(t *testing.T) {
		got := Welcome("Son", "German")
		want := "Hallo, Son"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Welcome("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

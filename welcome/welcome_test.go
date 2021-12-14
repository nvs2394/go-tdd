package main

import "testing"

func TestWelcome(t *testing.T) {
	got := Welcome("Son")
	want := "Hello, Son"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

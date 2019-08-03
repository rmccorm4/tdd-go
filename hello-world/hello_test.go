package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello with a given name", func(t *testing.T) {
		got := hello("Ryan")
		want := "Hello, Ryan"

		if got != want {
			t.Errorf("got %q - want %q", got, want)
		}
	})

	t.Run("Say 'Hello, World' when an empty string given", func(t *testing.T) {
		got := hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q - want %q", got, want)
		}
	})
}

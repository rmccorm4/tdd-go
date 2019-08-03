package main

import "testing"

func TestHello(t *testing.T) {
	assertion := func(t *testing.T, got, want string) {
		// Error gives line number of calling function rather than this function
		t.Helper()
		if got != want {
			t.Errorf("got %q - want %q", got, want)
		}
	}
	t.Run("Saying hello with a given name", func(t *testing.T) {
		got := hello("Ryan", "English")
		want := "Hello, Ryan"
		assertion(t, got, want)
	})

	t.Run("Say 'Hello, World' when an empty string given", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, World"
		assertion(t, got, want)
	})

	t.Run("Saying hello with a given name in Spanish", func(t *testing.T) {
		got := hello("Ryan", "Spanish")
		want := "Hola, Ryan"
		assertion(t, got, want)
	})

	t.Run("Saying hello world with no given name in Spanish", func(t *testing.T) {
		got := hello("", "Spanish")
		want := "Hola, World"
		assertion(t, got, want)
	})

	t.Run("Saying hello with a given name in French", func(t *testing.T) {
		got := hello("Ryan", "French")
		want := "Bonjour, Ryan"
		assertion(t, got, want)
	})
}

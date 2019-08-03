package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("Ryan")
	want := "Hello, Ryan"

	if got != want {
		t.Errorf("got %q - want %q", got, want)
	}
}

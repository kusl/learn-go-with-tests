package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Kus", "English")
	want := "Hello, Kus"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestHelloAgain(t *testing.T) {
	assertCorrectMessage := func(got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Kushal", "English")
		want := "Hello, Kushal"
		assertCorrectMessage(got, want)
	})
	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(got, want)
	})
	t.Run("en Espanol", func(t *testing.T) {
		got := Hello("Dora", "Español")
		want := "Hola, Dora"
		assertCorrectMessage(got, want)
	})
	t.Run("en français", func(t *testing.T) {
		got := Hello("Katherine", "Français")
		want := "Bonjour, Katherine"
		assertCorrectMessage(got, want)
	})
}

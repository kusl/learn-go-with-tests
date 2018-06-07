package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Kus", "English")
	want := "Hello, Kus"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

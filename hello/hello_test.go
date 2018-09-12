package main

import (
	"testing"
)

type Assert struct {
	*testing.T
}

func (a *Assert) CorrectMessage(got, want string) {
	a.Helper()
	if got != want {
		a.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestHello(t *testing.T) {
	a := Assert{t}
	a.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		a.CorrectMessage(got, want)
	})
	a.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		a.CorrectMessage(got, want)
	})
	a.Run("in French", func(t *testing.T) {
		got := Hello("Ola", "French")
		want := "Bonjour, Ola"
		a.CorrectMessage(got, want)
	})
	a.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		a.CorrectMessage(got, want)
	})
}

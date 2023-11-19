package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "english")
		want := "Hello, Chris"
		
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run("Say hello 'Hello, World' when an empty string is entered", func(t *testing.T){
		got := Hello("","")
		want := "Hello, World"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})

	t.Run("Say hello in Spanish", func(t *testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
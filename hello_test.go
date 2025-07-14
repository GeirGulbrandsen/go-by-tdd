package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("says hello to people", func(t *testing.T) {
		got := Hello("Geir", "")
		want := "Hello, Geir!"
		assertCorrectGreeting(t, got, want)
	})
	t.Run("says hello to the world when given an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertCorrectGreeting(t, got, want)

	})
	t.Run("says Hola to Spanish people", func(t *testing.T) {
		got := Hello("Elodie", spanish)
		want := "Hola, Elodie!"
		assertCorrectGreeting(t, got, want)
	})

	t.Run("says Bonjour to French people", func(t *testing.T) {
		got := Hello("Lola", french)
		want := "Bonjour, Lola!"
		assertCorrectGreeting(t, got, want)
	})
	t.Run("says Hei to Norwegian people", func(t *testing.T) {
		got := Hello("Geir", norwegian)
		want := "Hei, Geir!"
		assertCorrectGreeting(t, got, want)
	})
	t.Run("says Hello in English if language is not recognized", func(t *testing.T) {
		got := Hello("Geir", "Unrecognized")
		want := "Hello, Geir!"
		assertCorrectGreeting(t, got, want)
	})
}

func assertCorrectGreeting(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Hello() = %q; want %q", got, want)
	}
}

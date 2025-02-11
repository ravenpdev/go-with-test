package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("return default greeting 'Hello, World' if name and language is empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorretMessage(t, got, want)
	})

	t.Run("greet people in english if language is empty string", func(t *testing.T) {
		got := Hello("Raven", "")
		want := "Hello, Raven"

		assertCorretMessage(t, got, want)
	})

	t.Run("greet people in spanish", func(t *testing.T) {
		got := Hello("Raven", "Spanish")
		want := "Holla, Raven"

		assertCorretMessage(t, got, want)
	})

	t.Run("greet people in french", func(t *testing.T) {
		got := Hello("Raven", "French")
		want := "Bonjour, Raven"

		assertCorretMessage(t, got, want)
	})
}

func assertCorretMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

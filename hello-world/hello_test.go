package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := sayHello("Rusty", "English")
		want := "Hello Rusty"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello world' by default", func(t *testing.T) {
		got := sayHello("", "English")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := sayHello("Juan Pedro", "Spanish")
		want := "Hola Juan Pedro"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Deutsch", func(t *testing.T) {
		got := sayHello("Hans Gruber", "Deutsch")
		want := "Hallo Hans Gruber"
		assertCorrectMessage(t, got, want)
	})
}

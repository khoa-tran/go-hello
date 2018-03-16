package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := hello("Elise", "Spanish")
		want := "Hola, Elise"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := hello("David", "French")
		want := "Bonjour, David"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Vietnamese", func(t *testing.T) {
		got := hello("David", "Vietnamese")
		want := "Xin chao, David"
		assertCorrectMessage(t, got, want)
	})
}

func TestGreetingPrefix(t *testing.T) {
	t.Run("default in English", func(t *testing.T) {
		got := greetingPrefix("")
		want := "Hello, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := greetingPrefix("Spanish")
		want := "Hola, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := greetingPrefix("French")
		want := "Bonjour, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Vietnames", func(t *testing.T) {
		got := greetingPrefix("Vietnamese")
		want := "Xin chao, "
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

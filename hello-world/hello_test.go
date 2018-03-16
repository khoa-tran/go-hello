package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elise", "Spanish")
		want := "Hola, Elise"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("David", "French")
		want := "Bonjour, David"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Vietnamese", func(t *testing.T) {
		got := Hello("David", "Vietnamese")
		want := "Xin chao, David"
		assertCorrectMessage(t, got, want)
	})
}

func TestGreetingPrefix(t *testing.T) {
	t.Run("default in English", func(t *testing.T) {
		got := GreetingPrefix("")
		want := "Hello, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := GreetingPrefix("Spanish")
		want := "Hola, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := GreetingPrefix("French")
		want := "Bonjour, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Vietnames", func(t *testing.T) {
		got := GreetingPrefix("Vietnamese")
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

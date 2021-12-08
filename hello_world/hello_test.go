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
		got := Hello("cam")
		want := "hello, cam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'hello, world' when empty string is given", func(t *testing.T) {
		got := Hello("")
		want := "hello, world"
		assertCorrectMessage(t, got, want)
	})
}

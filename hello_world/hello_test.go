package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("cam")
	want := "hello, cam"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Cam")

	got := buffer.String()
	want := "Hello, Cam"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

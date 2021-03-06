package main

import (
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
	// 	buffer := &bytes.Buffer{}
	// 	spySleeper := &SpySleeper{}
	// 	Countdown(buffer, spySleeper)

	// 	got := buffer.String()
	// 	want := `3
	// 2
	// 1
	// Go!`

	// 	if got != want {
	// 		t.Errorf("got %q want %q", got, want)
	// 	}

	// 	if spySleeper.Calls != 4 {
	// 		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	// 	}
}

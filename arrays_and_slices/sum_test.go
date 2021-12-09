package main

import "testing"

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, want, got int) {
		t.Helper()
		if want != got {
			t.Errorf("expected %q but got %q", want, got)
		}
	}

	t.Run("Test 1", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectMessage(t, want, got)
	})

	t.Run("Test 2", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, want, got)
	})

}

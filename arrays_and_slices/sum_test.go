package main

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

	t.Run("large example", func(t *testing.T) {
		got := SumAllTails([]int{2, 5, 9}, []int{9, 1, 2, 3, 4, 5, 6, 7})
		want := []int{14, 28}
		checkSums(t, got, want)
	})

}

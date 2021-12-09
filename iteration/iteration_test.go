package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, repeated string) {
		t.Helper()
		if expected != repeated {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}
	t.Run("0 provided", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("-20 provided", func(t *testing.T) {
		repeated := Repeat("b", -20)
		expected := "bbbbb"
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("10 provided", func(t *testing.T) {
		repeated := Repeat("c", 10)
		expected := "cccccccccc"
		assertCorrectMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 20)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("c", 6))
	// Output: cccccc
}

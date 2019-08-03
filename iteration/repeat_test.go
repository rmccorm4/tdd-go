package repeat

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat a char", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %s - want %s", got, want)
		}
	})

	t.Run("Repeat a string", func(t *testing.T) {
		got := Repeat("abc", 5)
		want := "abcabcabcabcabc"

		if got != want {
			t.Errorf("got %s - want %s", got, want)
		}
	})

	t.Run("Repeat empty string", func(t *testing.T) {
		got := Repeat("", 5)
		want := ""

		if got != want {
			t.Errorf("got %s - want %s", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("double", 2)
	}
}

func ExampleRepeat() {
	x := Repeat("yes", 3)
	fmt.Println(x)
	// Output: yesyesyes
}

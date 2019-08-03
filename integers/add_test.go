package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	// t.Fatal("not implemented")
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected %d, but got %d", expected, sum)
	}
}

// ExampleX functions also get ran with `go test`, see `go test -v`
// ** But they will only run if you have a // Output: ... line at the end
func ExampleAdd() {
	x, y := 1, 5
	sum := Add(x, y)
	fmt.Printf("%d + %d = %d", x, y, sum)
	// Output: 1 + 5 = 6
}

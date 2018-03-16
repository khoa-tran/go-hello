package adder

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Add two similar numbers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("Add two different numbers", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

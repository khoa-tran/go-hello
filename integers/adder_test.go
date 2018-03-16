package adder

import "testing"

func TestAdder(t *testing.T) {
	t.Run("Add two similar numbers", func(t *testing.T) {
		sum := adder(2, 2)
		expected := 4
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("Add two different numbers", func(t *testing.T) {
		sum := adder(2, 3)
		expected := 5
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}

package repeat

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("should return empty", func(t *testing.T) {
		repeated := repeat("a", 0)
		expected := ""
		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	})

	t.Run("should repeat 2 times", func(t *testing.T) {
		repeated := repeat("a", 2)
		expected := "aa"
		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	})

	t.Run("should repeat 5 times", func(t *testing.T) {
		repeated := repeat("a", 5)
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	})
}

package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		checkSum(t, got, want, numbers)
	})

	t.Run("empty collection", func(t *testing.T) {
		numbers := []int{}
		got := Sum(numbers)
		want := 0
		checkSum(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("collection of one slice", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func checkSums(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func checkSum(t *testing.T, got, want int, numbers []int) {
	t.Helper()
	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

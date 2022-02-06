package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertHelperForInts := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	assertHelperForSlices := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("collection of 5 nums", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertHelperForInts(t, got, want)
	})

	t.Run("collection of 3 nums", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertHelperForInts(t, got, want)
	})

	t.Run("sum all slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assertHelperForSlices(t, got, want)
	})
}

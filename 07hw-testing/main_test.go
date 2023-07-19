package main

import (
	"sort"
	"testing"
)

func TestSortInts(t *testing.T) {
	s := []int{3, 5, 1, 0, -3}
	want := []int{-3, 0, 1, 3, 5}

	sort.Ints(s)

	if len(s) != len(want) {
		t.Fatalf("получилик оличество элементов: %d, ожидалось количество элементов: %d", len(s), len(want))
	}

	for i, n := range s {
		if n != want[i] {
			t.Fatalf("получили %d, ожидалось %d", n, want[i])
		}
	}
}

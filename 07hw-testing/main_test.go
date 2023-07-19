package main

import (
	"sort"
	"testing"
)

func TestSortInts(t *testing.T) {
	s := []int{3, 5, 1, 0, -3}
	want := []int{-3, 0, 1, 3, 5}

	sort.Ints(s)

	for i, n := range s {
		if n != want[i] {
			t.Fatalf("получили %d, ожидалось %d", n, want[i])
		}
	}
}

package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
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

func TestSortStrings(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		want []string
	}{
		{
			name: "test 1",
			s:    []string{"ax", "9", "0", "", "ab", "x"},
			want: []string{"", "0", "9", "ab", "ax", "x"},
		},
		{
			name: "test 2",
			s:    []string{"аз", "9", "0", "аб", "в", ""},
			want: []string{"", "0", "9", "аб", "аз", "в"},
		},
		{
			name: "test 3",
			s:    []string{},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Strings(tt.s)

			for i, s := range tt.s {
				if s != tt.want[i] {
					t.Errorf("получили %s, ожидалось %s", s, tt.want[i])
				}
			}
		})
	}
}

func sampleIntData() []int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	data := make([]int, 1_000_000)

	for i := 0; i < 1_000_000; i++ {
		data[i] = r.Intn(1000)
	}

	return data
}

package main

import (
	"reflect"
	"testing"
)

func Test_older(t *testing.T) {
	type args struct {
		a []ager
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{[]ager{&Employee{48}, &Customer{19}, &Customer{19}, &Employee{18}}},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := older(tt.args.a...); got != tt.want {
				t.Errorf("older() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_olderObj(t *testing.T) {
	type args struct {
		s []any
	}
	e2 := Employee{age: 48}
	c1 := Customer{age: 20}
	e1 := Employee{age: 48}
	c2 := Customer{age: 100}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "1",
			args: args{[]any{&e1, &c1, &c2, &e2}},
			want: &c2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := olderObj(tt.args.s...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("olderObj() = %v, want %v", got, tt.want)
			}
		})
	}
}

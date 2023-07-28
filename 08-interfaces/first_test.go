package homework

import "testing"

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
			args: args{[]ager{&Employee{18}, &Customer{19}}},
			want: 19,
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

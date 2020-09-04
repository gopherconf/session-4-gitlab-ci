package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		items []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Zero",
			args: args{
				items: []int64{0,0,0},
			},
			want: 0,
		},
		{
			name: "More than 3 items",
			args: args{
				items: []int64{1,10,30},
			},
			want: 41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.items...); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
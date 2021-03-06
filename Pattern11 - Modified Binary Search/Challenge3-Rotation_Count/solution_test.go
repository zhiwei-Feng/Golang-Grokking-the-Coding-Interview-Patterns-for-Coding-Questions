package challenge3rotationcount

import (
	"testing"
)

func Test_countRotations(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{10, 15, 1, 3, 8}}, 2},
		{"case2", args{[]int{4, 5, 7, 9, 10, -1, 2}}, 5},
		{"case3", args{[]int{1, 3, 8, 10}}, 0},
		{"case4", args{[]int{10, 1, 3, 8, 9}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRotations(tt.args.arr); got != tt.want {
				t.Errorf("countRotations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countRotations2(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{10, 15, 1, 3, 8}}, 2},
		{"case2", args{[]int{4, 5, 7, 9, 10, -1, 2}}, 5},
		{"case3", args{[]int{1, 3, 8, 10}}, 0},
		{"case4", args{[]int{10, 1, 3, 8, 9}}, 1},
		{"case5", args{[]int{3, 3, 7, 3}}, 3},
		{"case6", args{[]int{3, 3, 3, 3, 7}}, 0},
		{"case7", args{[]int{10, 1, 10, 10, 10}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRotations2(tt.args.arr); got != tt.want {
				t.Errorf("countRotations2() = %v, want %v", got, tt.want)
			}
		})
	}
}

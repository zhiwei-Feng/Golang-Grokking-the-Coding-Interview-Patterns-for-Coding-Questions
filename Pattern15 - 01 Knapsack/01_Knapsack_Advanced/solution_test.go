package knapsackadvanced

import (
	"testing"
)

func Test_solveKnapsack2(t *testing.T) {
	type args struct {
		profits  []int
		weights  []int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 7}, 22},
		{"case2", args{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 6}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveKnapsack2(tt.args.profits, tt.args.weights, tt.args.capacity); got != tt.want {
				t.Errorf("solveKnapsack2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveKnapsack(t *testing.T) {
	type args struct {
		profits  []int
		weights  []int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 7}, 22},
		{"case2", args{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 6}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveKnapsack(tt.args.profits, tt.args.weights, tt.args.capacity); got != tt.want {
				t.Errorf("solveKnapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}

package kthsmallestnumber

import "testing"

func Test_findKthSmallest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 5, 12, 2, 11, 5}, 3}, 5},
		{"case2", args{[]int{1, 5, 12, 2, 11, 5}, 4}, 5},
		{"case3", args{[]int{5, 12, 11, -1, 12}, 3}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthSmallest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}

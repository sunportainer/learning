package medium

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"leetcode209", args{7, []int{2, 3, 1, 2, 4, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSubArrayLenBetter(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"leetcode209", args{7, []int{2, 3, 1, 2, 4, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLenBetter(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLenBetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

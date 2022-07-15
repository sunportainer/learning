package dfs

import (
	"reflect"
	"testing"
)

func Test_canFinish2(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"leetcode210", args{2, [][]int{{1, 0}}}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish2(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canFinish2() = %v, want %v", got, tt.want)
			}
		})
	}
}

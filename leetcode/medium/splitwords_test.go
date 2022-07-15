package medium

import (
	"reflect"
	"testing"
)

func Test_partitionwithBack(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]string
	}{
		{"split workes", args{"aab"}, [][]string{{"a", "a", "b"}, {"aa", "b"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := partitionwithBack(tt.args.s); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("partition() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]string
	}{
		{"split workes", args{"aab"}, [][]string{{"a", "a", "b"}, {"aa", "b"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := partition(tt.args.s); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("partition() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

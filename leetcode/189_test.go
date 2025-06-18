package leetcode

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "Test 1",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "Test 2",
			nums: []int{-1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, -1, -100},
		},
		{
			name: "Test 3",
			nums: []int{5, 6, 5},
			k:    0,
			want: []int{5, 6, 5},
		},
		{
			name: "Test 4",
			nums: []int{5},
			k:    10,
			want: []int{5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.nums, tt.k)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}

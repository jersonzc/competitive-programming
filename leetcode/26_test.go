package leetcode

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Test 1",
			nums: []int{1, 1, 2},
			want: []int{1, 2},
		},
		{
			name: "Test 2",
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			assertInts(t, got, len(tt.want))
			assertKElements(t, got, tt.nums, tt.want)
		})
	}
}

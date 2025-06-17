package leetcode

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test 1",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "Test 2",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.nums)
			assertInts(t, got, tt.want)
		})
	}
}

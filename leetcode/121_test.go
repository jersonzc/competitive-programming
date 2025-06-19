package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test 1",
			nums: []int{7, 1, 5, 3, 6, 4},
			want: 5,
		},
		{
			name: "Test 2",
			nums: []int{7, 6, 4, 3, 1},
			want: 0,
		},
		{
			name: "Test 3",
			nums: []int{7},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.nums)
			assertInts(t, got, tt.want)
		})
	}
}

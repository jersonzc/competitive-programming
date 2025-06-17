package leetcode

import "testing"

func TestRemoveDuplicates2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Test 1",
			nums: []int{1, 1, 1, 2, 2, 3},
			want: []int{1, 1, 2, 2, 3},
		},
		{
			name: "Test 2",
			nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			want: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			name: "Test 3",
			nums: []int{5},
			want: []int{5},
		},
		{
			name: "Test 4",
			nums: []int{5, 5},
			want: []int{5, 5},
		},
		{
			name: "Test 5",
			nums: []int{5, 6},
			want: []int{5, 6},
		},
		{
			name: "Test 6",
			nums: []int{5, 5, 6},
			want: []int{5, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates2(tt.nums)
			assertInts(t, got, len(tt.want))
			assertKElements(t, got, tt.nums, tt.want)
		})
	}
}

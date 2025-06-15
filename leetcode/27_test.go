package leetcode

import (
	"sort"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		val          int
		want         int
		wantElements []int
	}{
		{
			name:         "Test 1",
			nums:         []int{3, 2, 2, 3},
			val:          3,
			want:         2,
			wantElements: []int{2, 2},
		},
		{
			name:         "Test 2",
			nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:          2,
			want:         5,
			wantElements: []int{0, 1, 3, 0, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.nums, tt.val)
			assertInts(t, got, tt.want)
			assertKElements(t, tt.want, tt.nums, tt.wantElements)
		})
	}
}

func assertInts(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertKElements(t testing.TB, k int, got, want []int) {
	t.Helper()
	sort.Ints(got[:k])
	sort.Ints(want)
	for i := 0; i < k; i++ {
		if got[i] != want[i] {
			t.Errorf("got %v want %v", got[:k], want)
			break
		}
	}
}

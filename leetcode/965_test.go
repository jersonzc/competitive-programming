package leetcode

import "testing"

func TestIsUnivalTree(t *testing.T) {
	tests := []struct {
		name string
		nums []*int
		want bool
	}{
		{
			name: "Test 1",
			nums: []*int{pointerToInt(1), pointerToInt(1), pointerToInt(1), pointerToInt(1), pointerToInt(1), nil, pointerToInt(1)},
			want: true,
		},
		{
			name: "Test 2",
			nums: []*int{pointerToInt(2), pointerToInt(2), pointerToInt(2), pointerToInt(5), pointerToInt(2)},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := buildTree(tt.nums)
			got := isUnivalTree(tree)
			assertBooleans(t, got, tt.want)
		})
	}
}

func assertBooleans(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got '%t' want '%t'", got, want)
	}
}

func pointerToInt(x int) *int { return &x }

func buildTree(values []*int) *TreeNode {
	n := len(values)

	nodes := make([]*TreeNode, n)
	for i, v := range values {
		if v != nil {
			nodes[i] = &TreeNode{Val: *v}
		}
	}

	for i := 0; i < n; i++ {
		if nodes[i] == nil {
			continue
		}

		leftIdx := 2*i + 1
		rightIdx := 2*i + 2

		if leftIdx < n {
			nodes[i].Left = nodes[leftIdx]
		}
		if rightIdx < n {
			nodes[i].Right = nodes[rightIdx]
		}
	}

	return nodes[0]
}

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrderTraversal(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left, nums)
	*nums = append(*nums, root.Val)
	inOrderTraversal(root.Right, nums)
}

func isUnivalTree(root *TreeNode) bool {
	var nums []int
	inOrderTraversal(root, &nums)

	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != curr {
			return false
		}
	}

	return true
}

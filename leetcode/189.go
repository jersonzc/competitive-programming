package leetcode

func rotate(nums []int, k int) {
	n := len(nums)

	if n == 1 {
		return
	}

	if k == 0 {
		return
	}

	r := k % n

	if r == 0 {
		return
	}

	temp := make([]int, n)
	for i := 0; i < r; i++ {
		temp[i] = nums[n-r+i]
	}
	for i := 0; i < n-r; i++ {
		temp[r+i] = nums[i]
	}

	for i := 0; i < n; i++ {
		nums[i] = temp[i]
	}
}

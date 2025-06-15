package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	curr := nums[0]
	unique := []int{curr}
	for i := 1; i < len(nums); i++ {
		if nums[i] != curr {
			curr = nums[i]
			unique = append(unique, curr)
		}
	}

	for i := 0; i < len(unique); i++ {
		nums[i] = unique[i]
	}

	return len(unique)
}

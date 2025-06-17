package leetcode

func removeDuplicates2(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	count := 1
	curr := nums[0]
	unique := []int{curr}
	for i := 1; i < len(nums); i++ {
		if nums[i] != curr {
			curr = nums[i]
			unique = append(unique, curr)
			count = 1
		} else {
			count++
			if count <= 2 {
				unique = append(unique, curr)
			}
		}
	}

	for i := 0; i < len(unique); i++ {
		nums[i] = unique[i]
	}

	return len(unique)
}

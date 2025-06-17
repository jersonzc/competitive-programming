package leetcode

import "sort"

func majorityElement(nums []int) int {
	n := len(nums)
	q := n / 2

	if n <= 2 {
		return nums[0]
	}

	sort.Ints(nums)

	curr := nums[0]
	count := 1
	for i := 1; i < n; i++ {
		if nums[i] != curr {
			if count > q {
				break
			}
			curr = nums[i]
			count = 1
		} else {
			count++
		}
	}

	return curr
}

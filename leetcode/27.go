package leetcode

func removeElement(nums []int, val int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		if nums[0] == val {
			nums = []int{}
			return 0
		} else {
			return 1
		}
	}

	c := 0

	i := 0
	j := n - 1

	for {
		if i > j {
			break
		}

		if i == j {
			if nums[i] == val {
				c++
			}
			break
		}

		if nums[i] == val {
			c++
			for {
				if i == j {
					return n - c
				}
				if nums[j] == val {
					c++
				} else {
					j--
					break
				}
				j--
			}
			nums[i], nums[j+1] = nums[j+1], nums[i]
		}

		i++
	}

	return n - c
}

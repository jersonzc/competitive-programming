package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	// ni tracks current value in nums2
	ni := 0
	v2 := 0
	if ni < n {
		v2 = nums2[ni]
	} else {
		return
	}

	// iterate over the nums1 array and allocate nums2 current values.
	it := 0
	for it < m+n {
		if v2 < nums1[it] {
			for j := m + n - 1; j > it; j-- {
				nums1[j] = nums1[j-1]
			}
			nums1[it] = v2
			ni++
			if ni < n {
				v2 = nums2[ni]
			} else {
				break
			}
		}
		it++
	}

	// allocate nums2 array values that were greater than any value found in nums1.
	for ni < n {
		nums1[m+ni] = nums2[ni]
		ni++
	}
}

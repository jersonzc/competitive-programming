package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	mi := m - 1
	ni := n - 1
	ki := m + n - 1

	for ki >= 0 {
		if mi < 0 {
			nums1[ki] = nums2[ni]
			ni--
			ki--
			continue
		}
		if ni < 0 {
			nums1[ki] = nums1[mi]
			mi--
			ki--
			continue
		}
		v1 := nums1[mi]
		v2 := nums2[ni]
		if v1 > v2 {
			nums1[ki] = v1
			mi--
		} else {
			nums1[ki] = v2
			ni--
		}
		ki--
	}
}

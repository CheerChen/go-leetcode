package arr

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	p := m + n - 1
	for p >= 0 && i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[p] = nums1[i]
			i--
		} else {
			nums1[p] = nums2[j]
			j--
		}
		p--
	}

	for p >= 0 && i >= 0 {
		nums1[p] = nums1[i]
		i--
		p--
	}
	for p >= 0 && j >= 0 {
		nums1[p] = nums2[j]
		j--
		p--
	}
}

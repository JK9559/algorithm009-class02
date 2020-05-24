package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 倒着填入元素
	p, q := m-1, n-1
	r := m + n - 1
	for p >= 0 && q >= 0 {
		if nums1[p] > nums2[q] {
			nums1[r] = nums1[p]
			p--
		} else {
			nums1[r] = nums2[q]
			q--
		}
		r--
	}
	// 如果nums1遍历完毕，需要把nums2的值填入，长度为0-q包含q
	if p < 0 {
		for i := 0; i <= q; i++ {
			nums1[i] = nums2[i]
		}
	}
}

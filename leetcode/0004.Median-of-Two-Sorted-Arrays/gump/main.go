package main

import (
	"fmt"
)

type para struct {
	nums1 []int
	nums2 []int
}

type inputAndOutput struct {
	para
	result float64
}

func main() {
	qs := []inputAndOutput{
		{
			para{[]int{1, 3}, []int{2, 4}},
			2.5,
		},
		{
			para{[]int{1, 2}, []int{3, 4}},
			2.5,
		},
		{
			para{[]int{5}, []int{}},
			5,
		},
		{
			para{[]int{1, 3}, []int{2}},
			2,
		},
	}

	if len(qs) == 0 {
		fmt.Println(0)
		return
	}

	for _, q := range qs {
		result := findMedianSortedArrays(q.para.nums1, q.para.nums2)
		if result != q.result {
			fmt.Printf("Error: expected %f, actual %f\n", result, q.result)
		} else {
			fmt.Printf("OK: expected %f, actual %f\n", result, q.result)
		}
	}
	return
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}

	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1
		} else {
			// 找到合适地划分了，需要输出最终结果了
			// 分为奇数偶数 2 中情况
			break
		}
	}

	fmt.Printf("nums = %v, mid = %d, nums2 = %v mid = %d\n", nums1, nums1Mid, nums2, nums2Mid)
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}

	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}

	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

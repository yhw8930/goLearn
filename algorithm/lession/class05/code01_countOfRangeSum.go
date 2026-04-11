package main

// https://leetcode.com/problems/count-of-range-sum/ 327:区间和的个数
// 给你一个整数数组 nums 以及两个整数 lower 和 upper 。求数组中，值位于范围 [lower, upper]
// （包含 lower 和 upper）之内的 区间和的个数 。
// 区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。
func countRangeSum(nums []int, lower int, upper int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	return process(sum, lower, upper, 0, len(sum)-1)
}

func process(sum []int, lower, upper, l, r int) int {
	if l == r {
		if sum[l] >= lower && sum[r] <= upper {
			return 1
		}
		return 0
	}
	m := l + (r-l)>>1
	return process(sum, lower, upper, l, m) + process(sum, lower, upper, m+1, r) + merge(sum, lower, upper, l, m, r)
}

func merge(sum []int, lower, upper, l, m, r int) int {
	ans := 0
	windowL := l
	windowR := l
	for i := m + 1; i <= r; i++ {
		minNum := sum[i] - upper
		maxNum := sum[i] - lower
		for windowR <= m && sum[windowR] <= maxNum {
			windowR++
		}
		for windowL <= m && sum[windowL] < minNum {
			windowL++
		}
		ans += windowR - windowL
	}

	help := make([]int, r-l+1)
	i := 0
	p1 := l
	p2 := m + 1
	for p1 <= m && p2 <= r {
		if sum[p1] <= sum[p2] {
			help[i] = sum[p1]
			p1++
		} else {
			help[i] = sum[p2]
			p2++
		}
		i++
	}
	for p1 <= m {
		help[i] = sum[p1]
		p1++
		i++
	}
	for p2 <= r {
		help[i] = sum[p2]
		p2++
		i++
	}
	for j := 0; j < len(help); j++ {
		sum[l+j] = help[j]
	}
	return ans
}

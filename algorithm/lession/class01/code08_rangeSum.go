package main

import "fmt"

// 题目：给定数组 nums，需要多次快速查询任意区间 i..j 的累加和。
// 如果每次查询都遍历区间，单次查询要 O(N)。
// 预处理前缀和数组 prefix，其中 prefix[k] 表示 0..k 的累加和。
// 核心思路：区间 i..j 的和可以由 prefix[j] - prefix[i-1] 在 O(1) 时间得到。
// 时间复杂度：构建前缀和 O(N)，单次查询 O(1)。
// 空间复杂度：O(N)。

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(rangeSum(nums, 0, 4))
	fmt.Println(rangeSum(nums, 2, 4))

}

// 前缀和，构造数组sum[i], i~j之间的和=sum[j]-sum[i]
func rangeSum(nums []int, left int, right int) int {
	rangeArr := buildRangeArr(nums)
	fmt.Println(rangeArr)
	if left == 0 {
		return rangeArr[right]
	}
	return rangeArr[right] - rangeArr[left-1]
}

func buildRangeArr(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] + nums[i]
	}
	return res
}

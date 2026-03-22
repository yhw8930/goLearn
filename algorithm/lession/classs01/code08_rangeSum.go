package main

import "fmt"

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

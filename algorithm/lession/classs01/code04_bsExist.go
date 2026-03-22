package main

import "fmt"

// exist 判断有序数组中是否存在目标数字（二分查找）
// sortedArr: 升序排列的整数切片
// num: 要查找的目标数字
// 返回值: 存在返回true，不存在返回false
func exist(sortedArr []int, num int) bool {
	if len(sortedArr) == 0 {
		return false
	}

	L, R := 0, len(sortedArr)-1
	// 闭区间 [L, R] 遍历，处理所有元素直到区间为空
	for L <= R {
		mid := L + ((R - L) >> 1)
		if sortedArr[mid] == num {
			return true
		} else if sortedArr[mid] > num {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return false
}

// 测试用例
func main() {
	// 测试场景1：存在目标值
	arr1 := []int{1, 3, 5, 7, 9, 11}
	fmt.Println(exist(arr1, 5))  // 输出：true
	fmt.Println(exist(arr1, 11)) // 输出：true

	// 测试场景2：不存在目标值
	fmt.Println(exist(arr1, 4))  // 输出：false
	fmt.Println(exist(arr1, 12)) // 输出：false

	// 测试场景3：空切片/边界值
	fmt.Println(exist(nil, 1))      // 输出：false
	fmt.Println(exist([]int{}, 2))  // 输出：false
	fmt.Println(exist([]int{2}, 2)) // 输出：true
	fmt.Println(exist([]int{2}, 3)) // 输出：false
}

package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7
	fmt.Println(mostLeftNoLessNumIndex(arr, target))
}

// arr有序， 找出>=num最左的位置
func mostLeftNoLessNumIndex(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	left, right := 0, len(arr)-1
	index := -1
	for left <= right {
		mid := left + (right-left)>>1
		if arr[mid] >= target {
			right = mid - 1
			index = mid
		} else if arr[mid] < target {
			left = mid + 1
		}
	}
	return index
}

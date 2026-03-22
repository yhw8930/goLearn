package main

import (
	"fmt"
)

func main() {
	// 待排序的原始数组。
	original := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original Array ->", original)
	fmt.Println("---------------------------------")

	// 把排序函数放到同一张表，方便统一演示和对比输出。
	sorts := []struct {
		name string
		fn   func([]int)
	}{
		{"BubbleSort", BubbleSort},
		{"SelectionSort", SelectionSort},
		{"InsertionSort", InsertionSort},
		{"ShellSort", ShellSort},
		{"MergeSort", MergeSort},
		{"QuickSort", QuickSort},
		{"HeapSort", HeapSort},
	}

	for _, s := range sorts {
		// 每次复制一份原始数据，避免前一次排序结果影响后续算法。
		// 使用 append([]int(nil), original...) 是一个高效的切片复制方法。
		nums := append([]int(nil), original...)
		s.fn(nums)
		fmt.Printf("%-15s -> %v\n", s.name, nums)
	}
}

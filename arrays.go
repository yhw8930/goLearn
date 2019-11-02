package main

import "fmt"

func printArray(arr []int) {
	for key, value := range arr {
		fmt.Println(key, value)
	}
	arr[0] = 100
}
func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	for key, value := range arr3 {
		fmt.Println(key, value)
	}
	printArray(arr3[:])
	for key, value := range arr3 {
		fmt.Println(key, value)
	}
}

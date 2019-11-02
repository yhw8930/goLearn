package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}
func main() {
	fmt.Println("Creating slice")
	var s []int //Zero value for slice is nil
	for i := 0; i < 100; i++ {
		//printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 16, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Coping slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting slice")
	s4 := append(s2[:3], s2[4:]...)
	printSlice(s4)
	//slice 只会向后切,含头部时capcity会改变，后切不改变
	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)
}

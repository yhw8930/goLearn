package main

import "fmt"

func main() {
	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
	fmt.Printf("%T\n", &v)
	m := map[string]int{
		"sd":   1,
		"sdsd": 1,
	}
	for _, v := range m {
		fmt.Println(v)
	}

}

func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}

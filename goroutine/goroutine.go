package main

import (
	"fmt"
	"time"
)

func main() {
	//var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				//a[i]++
				//fmt.Println(a[i])
				//runtime.Gosched()
				fmt.Printf("Hello from goroutine %d\n", i)
				//go run -race file
			}
		}(i)
	}
	time.Sleep(time.Minute)
	//fmt.Println(a)
}

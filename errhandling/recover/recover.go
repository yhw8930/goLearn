package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()
	//panic(errors.New("this is an error"))
	//a := 0
	//b := 5 / a
	//fmt.Println(b)
	panic(123)
}
func main() {
	tryRecover()
}

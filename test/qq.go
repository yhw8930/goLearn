package main

import "fmt"

func main() {
	var i int64
	//var j int64
	if &i == nil {
		fmt.Println("11")
	} else {
		fmt.Println("22")
	}
}

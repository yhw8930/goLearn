package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.Atoi("123"))
	fmt.Println(strconv.Itoa(111))
	fmt.Println(strconv.ParseInt("12345", 10, 64))
	fmt.Println(strconv.FormatInt(1234567, 10))
}

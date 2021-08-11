package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var result1 interface{} = Ups(1)
	var result2 interface{} = Ups(2)
	var result3 interface{} = Ups(3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
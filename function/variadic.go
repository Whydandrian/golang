package main

import "fmt"

func sumAll(number ...int) int  {
	total := 0
	for _, value:= range number {
		total += value
	}
	return total
}

func main() {
	total := sumAll(12, 15, 21, 22, 42)
	fmt.Println(total)
	numbers := []int{12, 51, 23, 44}
	res := sumAll(numbers...)
	fmt.Println(res)
}
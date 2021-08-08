package main

func Sum(x, y int) int {
	return x + y
}

func SumVariadic(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
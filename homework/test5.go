package main

func add(args ...int) int  {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func main() {
	add(1, 2)
	var x = nil
	var s string = nil
	var xsx interface{} = nil
	var c error = nil
}
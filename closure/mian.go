package main

import "fmt"

func main() {
	name := "wahyudi"
	counter := 0

	increment := func() {
		name := "andrian"
		// dengan variable definition di bawah ini akan menimpa isi variable name
		// yang berada di atas counter
		// name = "wahyu"
		fmt.Println("increment")
		counter++
		fmt.Println(name)
	}

	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
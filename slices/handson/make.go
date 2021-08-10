package main

import "fmt"

func main() {
	// Menggunakan Make
	
	// Buat slice of int bernama "a" dengan panjang 3, {1, 2,3}
	a := make([]int, 3)
	a = []int{1, 2, 3}
	printSlice("a", a)
	
	// Buat slice of int bernama "b" dengan panjang 5 {1, 2, 3, 4, 5}
	b := make([]int, 5)
	b = []int{1, 2, 3, 4, 5}
	printSlice("b", b)

	// Buat variabel c dengan dua data awal a
	c := a[:2]
	printSlice("c", c)
	
	// Buat variabel d dengan data ke 2 sampai 4 variabel b
	d := b[1:3]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",s, len(x), cap(x), x)
}
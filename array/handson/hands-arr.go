package main

import "fmt"

func main() {
	// buat array a string dengan panjang 2
	var a [2]string
	// masukkan nilai pertama "Go"
	a[0] = "Go"
	// masukkan nilai kedua "C"
	a[1] = "C"
	// print nilai pertama, dan kedua
	fmt.Println(a[0], a[1])
	fmt.Println("nilai pertama", a[0])
	fmt.Println("nilai kedua", a[1])
	// print a
	fmt.Println(a)


	// buat variabel primes dan buatlah
	// menggunakan array literals dengan panjang 6
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// buat 6 angka bilangan prima
	// print primes
	fmt.Println(primes)

}

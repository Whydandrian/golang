package main

import "fmt"

func main() {
	// buat variabel primes dan buatlah
	// menggunakan slice literals dengan panjang 6
	// buat 6 angka bilangan prima
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// print primes
fmt.Println(primes)
// print primes dari data ke 3 sampai 5
fmt.Println(primes[2:5])


}
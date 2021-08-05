package main

import "fmt"

func main() {
	// Constant adalah variable yang isinya tidak dapat diubah lagi
	const firstName string = "Wahyudi"
	const lastName = "Andrian"
	const value = 100

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
	
	// Declare Multi Constant
	const (
		nama string = "Wahyudi Andrian"
		alamat = "Lumajang"
		umur = 24
	)
	
	fmt.Println(nama)
	fmt.Println(alamat)
	fmt.Println(umur)


}

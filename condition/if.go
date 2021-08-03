package main

import "fmt"

func main() {
	var name = "Wahyudi"

	if name == "Wahyudi" {
		fmt.Println("Hello Wahyudi")
	} else {
		fmt.Println("kamu  bukan wahyudi")
	}

	//Short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
package main

import "fmt"

func main() {
	name := "Wahyudi"

	switch name {
	case "Wahyudi":
		fmt.Println("Hai Wahyudi..")
	case "Budi":
		fmt.Println("Hai Joko..")
	default:
		fmt.Println("Hai boleh kenalan..")
		fmt.Println("Nama lengkap yaa")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}
}
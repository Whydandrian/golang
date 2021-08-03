package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// For range with slice
	slice := []string{"wahyudi", "andrian","ganteng"}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	for i, value := range slice {
		fmt.Println("Index ke ", i, "= ", value)
	}

	// jika kita ingin menampilkan data value saja
	// untuk index yang dideklarasikan bisa diubah
	// atau diganti dengan simbol underscore _
	// karena jika variable tidak di pakai di golang
	// akan terdeteksi error
	for _, value := range slice {
		fmt.Println(value)
	}

	person := make(map[string]string)

	person["name"] = "Wahyudi"
	person["role"] = "Software Dev."

	for key, value := range person{
		fmt.Println(key, "=", value)
	}
}
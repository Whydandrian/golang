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

}
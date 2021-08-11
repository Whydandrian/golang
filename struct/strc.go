package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var wahyudi Customer

	wahyudi.Name = "Wahyudi Andrian"
	wahyudi.Address = "Lumajang"
	wahyudi.Age = 24

	fmt.Println(wahyudi.Name)
	fmt.Println(wahyudi.Address)
	fmt.Println(wahyudi.Age)
}

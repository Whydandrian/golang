package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Married bool
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hi ", name, "My name is ", customer.Name)
}

func main() {
	var wahyudi Customer

	wahyudi.Name = "Wahyudi Andrian"
	wahyudi.Address = "Lumajang"
	wahyudi.Age = 24

	wahyudi.sayHi("admin")

	// fmt.Println(wahyudi.Name)
	// fmt.Println(wahyudi.Address)
	// fmt.Println(wahyudi.Age)

	// whyd := Customer {
	// 	Name: "whydandrian",
	// 	Address: "Lumajang-Jatim",
	// 	Age: 24,
	// }
	// fmt.Println(whyd)
	// kode di bawah ini akan error karena deklarasi value tidak sesuai dg
	// urutan field struct Customer
	// yudi := Customer{"Yuyud", "magelang", 22}
}

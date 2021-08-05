// Untuk membuat variable di awali : var variable_name

package main

import "fmt"

func main() {
	var name string
	var my_name = "Wahyudi Andrian"
	var age = 24

	name = "Wahyudi Andrian"
	fmt.Println(name)

	name = "Wahyudi"
	fmt.Println(name)
	fmt.Println("Name : ", my_name)
	fmt.Println("Age : ", age)

	//Simple Declaration Variable
	country := "Indonesia"

	fmt.Println(country)

	var (
		firstName = "Wahyudi"
		lastName = "Andrian"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}

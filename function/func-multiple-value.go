package main

import "fmt"

func getFullName() (string, string, string) {
	return "Wahyudi", "Andrian", "Ganteng"
}

func main() {
	// firstName, LastName, addFitures := getFullName()
	firstName, LastName, _ := getFullName()
	fmt.Println(firstName)
	fmt.Println(LastName)
}
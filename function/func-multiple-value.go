package main

import "fmt"

func getFullName() (string, string) {
	return "Wahyudi", "Andrian"
}

func main() {
	firstName, LastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(LastName)
}
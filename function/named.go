package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Wahyudi"
	middleName = "ganteng"
	lastName = "Andrian"

	return
}

func main() {
	firstName, middleName, lastName := getCompleteName()

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(middleName)
}
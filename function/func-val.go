package main

import "fmt"

func getGoodBye(name string) string  {
	return "Good bye " + name
}

func main() {
	sayGoodBye := getGoodBye
	result := sayGoodBye("Wahyudi")
	fmt.Println(result)
}
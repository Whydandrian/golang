package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil logging")
}

func runApplication(value int)  {

	defer	logging()

	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("hasil baginya :", result)
}

func main() {
	runApplication(0)
}
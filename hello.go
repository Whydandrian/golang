package main

import (
	"fmt"
	"os"
)

func main()  {
	// fmt.Println("Hi.. I'm Wahyudi Andrian")
	// fmt.Println("I'm from Lumajang, East Java")

	var s string
	sep := ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println((s))
}
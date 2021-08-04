package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		// Break for loop
		// if i == 5 {
		// 	break
		// }

		if i%2 == 0 {
			continue
		}
		fmt.Println("No. ", i)
	}
}
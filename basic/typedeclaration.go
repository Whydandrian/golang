package main

import "fmt"

func main() {

	// Type Declaration
	// memudahkan dalam pendeklarasian variabel

	type NomorKTP string
	type Married bool

	var nomorKTP NomorKTP = "1872348123873"
	fmt.Println(nomorKTP)

	var statusMarried Married = true
	fmt.Println(statusMarried)

}

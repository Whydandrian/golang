package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	// variable standard

	a := rand.Intn(56)
	fmt.Println("Nomor favorit saya adalah :", a)

	// declare variable in 1 row
	c, python, java := true, false, "no!"
	fmt.Println(c, python, java)

	// convert data type
	myNumber := 100
	num := strconv.Itoa(myNumber)
	fmt.Println("Hasil dari konvert data tipe integer ke string adalah : ", num)

	x, y := 3, 4
	result := math.Sqrt(float64((x*x + y*y)))
	fmt.Println(result)

	z := uint(result)

	println(x, y, z)

	g := "tuet"
	uui := 123.2
	in := 12

	fmt.Printf("Tipe data : %T\n", g)
	fmt.Printf("Tipe data : %T\n", uui)
	fmt.Printf("Tipe data : %T\n", in)
	
	// buatlah variabel p dengan pointer i
	ii, jj := 42, 2701
	p := &ii
	fmt.Println(*p)
	*p = 21
	fmt.Println(ii)
	p = &jj
	*p = *p / 37
	fmt.Println(jj)
}

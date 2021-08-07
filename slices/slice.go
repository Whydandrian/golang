package main

import "fmt"

func main() {
	// var fruits = []string{"apple", "grape", "banana", "melon", "tomato"}
	// var vegetables = []string{"spinach", "cucumber", "kale", "lettuce", "tomato?"}

	// fmt.Println(fruits[0])
	// fmt.Println(vegetables[1])

	// likedFruits := fruits[0:3]
	// likedVegetables := vegetables[2:4]
	// fmt.Println(likedFruits)
	// fmt.Println(likedVegetables)

	// Memecah string dan menambahkan per string ke dalam array
	var runes []rune
	for _, r := range "Hi, Wahyudi" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	a := make([]int, 5)
	// make([array]tuipe data, panjang_slice, kapasitas_slice)
	b := make([]int, 7, 5)

	// var x, y []int
	// for i := 0; i < 10; i++ {
	// 	y = append(x, i)
	// 	fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	// 	x = y
	// }

	// s := make([]string, 3)
  //   fmt.Println("emp:", s)

  //   s[0] = "a"
  //   s[1] = "b"
  //   s[2] = "c"
  //   fmt.Println("set:", s)
  //   fmt.Println("get:", s[2])

  //   fmt.Println("len:", len(s))

  //   s = append(s, "d")
  //   s = append(s, "e", "f")
  //   fmt.Println("apd:", s)

  //   c := make([]string, len(s))
  //   copy(c, s)
  //   fmt.Println("cpy:", c)

  //   l := s[2:5]
  //   fmt.Println("sl1:", l)

  //   l = s[:5]
  //   fmt.Println("sl2:", l)

  //   l = s[2:]
  //   fmt.Println("sl3:", l)

  //   t := []string{"g", "h", "i"}
  //   fmt.Println("dcl:", t)

  //   twoD := make([][]int, 3)
  //   for i := 0; i < 3; i++ {
  //       innerLen := i + 1
  //       twoD[i] = make([]int, innerLen)
  //       for j := 0; j < innerLen; j++ {
  //           twoD[i][j] = i + j
  //       }
  //   }
  //   fmt.Println("2d: ", twoD)


}
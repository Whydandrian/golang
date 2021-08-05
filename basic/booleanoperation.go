package main

import "fmt"

func main(){
	var nilaiakhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiakhir >= 80
	var nilaiAbsensi bool = absensi >= 80

	var lulus bool = lulusNilaiAkhir && nilaiAbsensi

	fmt.Println(lulus)
}

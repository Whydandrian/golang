package main

import "fmt"

func main() {
	// buat variabel m dengan map key: string dan value: int
	m := map[string]int{
		"Answer": 42,
	}

	// beri nilai "Answer" dengan 42
	fmt.Println("The value:", m["Answer"])

	// ganti nilai "Answer" dengan 48
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// hapus "Answer
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// gunakan pengecekan: elem, ok = m[key]
	v, ok := m["Answer"]



	fmt.Println("The value:", v, "Present?", ok)
}
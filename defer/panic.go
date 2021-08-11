package main

import "fmt"

func endApp() {
	fmt.Println("aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	// runApp(false)
}
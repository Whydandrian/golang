package main

import "fmt"

func endApp() {
	// add recover to get message while panic
	message := recover()
	if message != nil {
		fmt.Println("error dengan message: ", message)
	}
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
	runApp(false)
}
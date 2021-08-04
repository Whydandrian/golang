package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hai, boleh kenalan?"
	} else {
		return "Hai " + name
	}
}

func main() {
	res := getHello("")
	fmt.Println(res)
}
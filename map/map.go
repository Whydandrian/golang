package main

import "fmt"

func main()  {
	person := map[string]string{
		"name":"wahyudi",
		"address":"Lumajang",
	}
	person["title"] = "Backend Dev."

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Wahyudi"
	book["status"] = "released"

	fmt.Println(book)
	
	delete(book, "status")
	fmt.Println(book)
}
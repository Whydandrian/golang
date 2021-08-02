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
	
}
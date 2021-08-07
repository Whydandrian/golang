package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main()  {
	// person := map[string]string{
	// 	"name":"wahyudi",
	// 	"address":"Lumajang",
	// }
	// person["title"] = "Backend Dev."

	// fmt.Println(person)
	// fmt.Println(person["name"])
	// fmt.Println(person["address"])
	// fmt.Println(person["title"])

	// var book map[string]string = make(map[string]string)
	// book["title"] = "Belajar Go-Lang"
	// book["author"] = "Wahyudi"
	// book["status"] = "released"

	// fmt.Println(book)
	
	// delete(book, "status")
	// fmt.Println(book)
	fmt.Println(person{"Bob", 20})
	
	fmt.Println(person{name: "Alice", age: 30})
	
	fmt.Println(person{name: "Fred"})
	
	fmt.Println(&person{name: "Ann", age: 40})
	
	fmt.Println(newPerson("Jon"))
	
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	
	sp := &s
	fmt.Println(sp.age)
	
	sp.age = 51
	fmt.Println(sp.age)

}
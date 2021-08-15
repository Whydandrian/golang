// package main

// import "fmt"
// import "net/http"
// import "time"

// func handlerIndex(w http.ResponseWriter, r *http.Request) {
// 	var message = "Welcome"
// 	w.Write([]byte(message))
// }

// func handlerHello(w http.ResponseWriter, r *http.Request) {
// 	var message = "Hello world!"
// 	w.Write([]byte(message))
// }

// func main() {
// 	http.HandleFunc("/", handlerIndex)
// 	http.HandleFunc("/index", handlerIndex)
// 	http.HandleFunc("/hello", handlerHello)

// 	var address = "localhost:8000"
// 	fmt.Printf("server started at %s\n", address)

// 	server := new(http.Server)
// 	server.Addr = address
// 	server.ReadTimeout = time.Second * 10
// 	server.WriteTimeout = time.Second * 10

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }

package main

import "github.com/gofiber/fiber/v2"
func (c *fiber.Ctx) error  {
	return c.SendString("Hi.. Wahyudi!")
}
func main() {
	app := fiber.New()

	app.Get("/", index)

	app.Listen(":3000")
}
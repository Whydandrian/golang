package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://sameer-kumar-aztro-v1.p.rapidapi.com/?sign=libra&day=today"
	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("x-rapidapi-key", "")
	req.Header.Add("x-rapidapi-host", "sameer-kumar-aztro-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}

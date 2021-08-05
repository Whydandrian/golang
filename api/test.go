package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	url := "https://sameer-kumar-aztro-v1.p.rapidapi.com/?sign=libra&day=today"
	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("x-rapidapi-key", "ed5f837f73msh394956de50910e0p1908e5jsn80aa43a20b2c")
	req.Header.Add("x-rapidapi-host", "sameer-kumar-aztro-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}
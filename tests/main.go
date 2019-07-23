package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Post("http://localhost:8080/fizzbuzz?fizz=riri&buzz=loulou&val1=2&val2=5&limit=1000", "", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func String(r []byte) string {
	s := make([]string, 0, len(r))
	for _, i := range r {
		s = append(s, string(i))
	}
	return fmt.Sprintf(strings.Join(s, "."))
}

func main() {
	resp, err := http.Post("http://localhost:8080/fizzbuzz?fizz=riri&buzz=loulou&val1=2&val2=5&limit=100", "", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(String(body))
}

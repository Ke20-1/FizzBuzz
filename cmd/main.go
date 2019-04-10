package main

import (
	"log"
	"net/http"

	"gitlab.app.treezor.com/card/example/FizzBuzz/handlers"
)

func main() {
	http.HandleFunc("/fizzbuzz", handlers.FizzBuzzHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

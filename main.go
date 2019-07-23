package main

import (
	"net/http"
	"os"

	"FizzBuzz/handlers"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Out = os.Stdout
	log.Info("Starting server")

	http.HandleFunc("/fizzbuzz", handlers.FizzBuzzHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

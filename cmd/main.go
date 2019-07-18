package main

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"gitlab.app.treezor.com/card/example/FizzBuzz/handlers"
)

func main() {
	log := logrus.New()
	log.Out = os.Stdout
	log.Info("Starting server")

	http.HandleFunc("/fizzbuzz", handlers.FizzBuzzHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"net/http"
	"os"

	"fizzbuzz/handlers"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Out = os.Stdout
	log.Info("Starting server")

	http.HandleFunc("/fizzbuzz", handlers.FizzbuzzHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

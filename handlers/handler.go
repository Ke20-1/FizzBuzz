package handlers

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
	"gitlab.app.treezor.com/card/example/FizzBuzz/algo"
)

// LogMiddleware log info
// func LogMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//
// 		logger.WithFields(logger.Fields{
// 			"animal": "walrus",
// 		}).Info("A walrus appears")
// 		next.ServeHTTP(w, r)
// 	})
// }

var log = logrus.New()

// FizzBuzzHandler is entrypoint for return fizzBuzz request
func FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("incomming request")
	if r.Method == "POST" {
		fb, err := parseParameters(r.URL.Query())
		if err != nil {
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				log.Errorf("error to send error response : %v", err)
			}
		}
		buffer := &bytes.Buffer{}

		err = gob.NewEncoder(buffer).Encode(algo.FizzBuzzAlgo(fb))
		if err != nil {
			log.Errorf("error to encode response : %v", err)
		}
		log.Info(algo.FizzBuzzAlgo(fb))

		_, err = w.Write(buffer.Bytes())
		if err != nil {
			log.Errorf("error to sending response : %v", err)
		}
	} else {
		log.Info("request discard because not a post")
	}
}

// parseParameters receive argument from fizzbuzz request and return corrected argument for fizzbuzz algo
func parseParameters(param url.Values) (*algo.FizzBuzz, error) {
	fizz, okfizz := param["fizz"]
	buzz, okbuzz := param["buzz"]
	val1, okval1 := param["val1"]
	val2, okval2 := param["val2"]
	limit, oklimit := param["limit"]

	if !okfizz || len(fizz[0]) < 1 {
		return nil, fmt.Errorf("url Param 'key' is missing")
	}
	if !okbuzz || len(buzz[0]) < 1 {
		log.Println("url Param 'key' is missing")
		return nil, fmt.Errorf("url Param 'key' is missing")
	}
	if !okval1 || len(val1[0]) < 1 {
		log.Println("url Param 'key' is missing")
		return nil, fmt.Errorf("url Param 'key' is missing")
	}
	if !okval2 || len(val2[0]) < 1 {
		log.Println("url Param 'key' is missing")
		return nil, fmt.Errorf("url Param 'key' is missing")
	}
	if !oklimit || len(limit[0]) < 1 {
		log.Println("url Param 'key' is missing")
		return nil, fmt.Errorf("url Param 'key' is missing")
	}

	fizzF := fizz[0]
	buzzF := buzz[0]
	val1F := val1[0]
	val2F := val2[0]
	limitF := limit[0]

	log.Println("Url Param 'fizz' is: " + fizzF)
	log.Println("Url Param 'buzzF' is: " + buzzF)
	log.Println("Url Param 'val1F' is: " + val1F)
	log.Println("Url Param 'val2F' is: " + val2F)
	log.Println("Url Param 'limitF' is: " + limitF)
	val1Num, err := strconv.Atoi(val1F)
	if err != nil {
		return nil, err
	}
	val2Num, err := strconv.Atoi(val2F)
	if err != nil {
		return nil, err
	}
	limitNum, err := strconv.Atoi(limitF)
	if err != nil {
		return nil, err
	}
	return &algo.FizzBuzz{
		Fizz:  fizzF,
		Buzz:  buzzF,
		Int1:  val1Num,
		Int2:  val2Num,
		Limit: limitNum,
	}, nil
}

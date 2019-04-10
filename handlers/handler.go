package handlers

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"gitlab.app.treezor.com/card/example/FizzBuzz/algo"
)

// FizzBuzzHandler is entrypoint for return fizzBuzz request
func FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fizzF, buzzF, int1Num, int2Num, limitNum := parseParameters(r.URL.Query())
		buffer := &bytes.Buffer{}
		fb := algo.FizzBuzz{
			Fizz:  fizzF,
			Buzz:  buzzF,
			Int1:  int1Num,
			Int2:  int2Num,
			Limit: limitNum,
		}
		gob.NewEncoder(buffer).Encode(algo.FizzBuzzAlgo(fb))
		fmt.Println(algo.FizzBuzzAlgo(fb))

		w.Write(buffer.Bytes())
	}
}

// parseParameters receive argument from fizzbuzz request and return corrected argument for fizzbuzz algo
func parseParameters(param url.Values) (fizzF, buzzF string, int1Num, int2Num, limitNum int) {
	fizz, okfizz := param["fizz"]
	buzz, okbuzz := param["buzz"]
	int1, okint1 := param["int1"]
	int2, okint2 := param["int2"]
	limit, oklimit := param["limit"]

	if !okfizz || len(fizz[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
	if !okbuzz || len(buzz[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
	if !okint1 || len(int1[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
	if !okint2 || len(int2[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
	if !oklimit || len(limit[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	fizzF = fizz[0]
	buzzF = buzz[0]
	int1F := int1[0]
	int2F := int2[0]
	limitF := limit[0]

	log.Println("Url Param 'fizz' is: " + string(fizzF))
	log.Println("Url Param 'buzzF' is: " + string(buzzF))
	log.Println("Url Param 'int1F' is: " + string(int1F))
	log.Println("Url Param 'int2F' is: " + string(int2F))
	log.Println("Url Param 'limitF' is: " + string(limitF))
	int1Num, err := strconv.Atoi(string(int1F))
	if err != nil {

		log.Print("error value receive int1 ", err)
	}
	int2Num, err = strconv.Atoi(string(int2F))
	if err != nil {
		log.Print("error value receive int2 ", err)
	}
	limitNum, err = strconv.Atoi(string(limitF))
	if err != nil {
		log.Print("error value receive int2 ", err)
	}
	return
}

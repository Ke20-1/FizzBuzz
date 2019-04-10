package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"log"
	"net/http"
	"net/url"
	"strconv"
)

func ParseParameters(param url.Values) (fizzF, buzzF string, int1Num, int2Num, limitNum int) {
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

func fizzBuzz(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fizzF, buzzF, int1Num, int2Num, limitNum := ParseParameters(r.URL.Query())
		buffer := &bytes.Buffer{}
		gob.NewEncoder(buffer).Encode(FizzBuzzAlgo(fizzF, buzzF, int1Num, int2Num, limitNum))
		fmt.Println(FizzBuzzAlgo(fizzF, buzzF, int1Num, int2Num, limitNum))

		w.Write(buffer.Bytes())
	}
}

func FizzBuzzAlgo(fizz, buzz string, int1, int2, limit int) (resultFizzBuzz []string) {
	resultFizzBuzz = append(resultFizzBuzz, ",")

	for i := 0; i <= limit; i++ {
		if i == 0 {
			resultFizzBuzz = append(resultFizzBuzz, strconv.Itoa(i))
			resultFizzBuzz = append(resultFizzBuzz, ",")
			continue
		}
		if (i%int1 == 0) && (i%int2 == 0) {
			resultFizzBuzz = append(resultFizzBuzz, fizz+buzz)
			resultFizzBuzz = append(resultFizzBuzz, ",")
			continue
		}
		if i%int1 == 0 {
			resultFizzBuzz = append(resultFizzBuzz, fizz)
			resultFizzBuzz = append(resultFizzBuzz, ",")
			continue
		}
		if i%int2 == 0 {
			resultFizzBuzz = append(resultFizzBuzz, buzz)
			resultFizzBuzz = append(resultFizzBuzz, ",")
			continue
		}
		resultFizzBuzz = append(resultFizzBuzz, strconv.Itoa(i))
		resultFizzBuzz = append(resultFizzBuzz, ",")
	}

	return
}

func main() {

	http.HandleFunc("/fizzbuzz", fizzBuzz)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

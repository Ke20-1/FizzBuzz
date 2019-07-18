package algo

import "strconv"

// FizzBuzz is struct for make algo
type FizzBuzz struct {
	Fizz  string
	Buzz  string
	Int1  int
	Int2  int
	Limit int
}

// FizzBuzzAlgo is mainy algo
func FizzBuzzAlgo(fizzBuzz *FizzBuzz) (resultFizzBuzz string) {
	//resultFizzBuzz = append(resultFizzBuzz, ",")

	for i := 0; i <= fizzBuzz.Limit; i++ {
		if i == 0 {
			resultFizzBuzz += strconv.Itoa(i) + ","
			// resultFizzBuzz += "," + ","
			continue
		}
		if (i%fizzBuzz.Int1 == 0) && (i%fizzBuzz.Int2 == 0) {
			resultFizzBuzz += fizzBuzz.Fizz + fizzBuzz.Buzz + ","
			// resultFizzBuzz += "," + ","
			continue
		}
		if i%fizzBuzz.Int1 == 0 {
			resultFizzBuzz += fizzBuzz.Fizz + ","
			// resultFizzBuzz += "," + ","
			continue
		}
		if i%fizzBuzz.Int2 == 0 {
			resultFizzBuzz += fizzBuzz.Buzz + ","
			// resultFizzBuzz += "," + ","
			continue
		}
		resultFizzBuzz += strconv.Itoa(i) + ","
		// resultFizzBuzz = append(resultFizzBuzz, ",")
	}
	return resultFizzBuzz[:len(resultFizzBuzz)-1]
}

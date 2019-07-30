package algo

import "strconv"

// Fizzbuzz is struct for make algo
type Fizzbuzz struct {
	Fizz  string
	Buzz  string
	Int1  int
	Int2  int
	Limit int
}

// FizzbuzzAlgo is mainy algo
func FizzbuzzAlgo(fizzbuzz *Fizzbuzz) (resultfizzbuzz string) {
	for i := 0; i <= fizzbuzz.Limit; i++ {
		if i%(fizzbuzz.Int1*fizzbuzz.Int2) == 0 {
			resultfizzbuzz += fizzbuzz.Fizz + fizzbuzz.Buzz + ","
			continue
		}
		if i == 0 {
			resultfizzbuzz += strconv.Itoa(i) + ","
			continue
		}
		if i%fizzbuzz.Int1 == 0 {
			resultfizzbuzz += fizzbuzz.Fizz + ","
			continue
		}
		if i%fizzbuzz.Int2 == 0 {
			resultfizzbuzz += fizzbuzz.Buzz + ","
			continue
		}
		resultfizzbuzz += strconv.Itoa(i) + ","
	}
	return resultfizzbuzz[:len(resultfizzbuzz)-1]
}

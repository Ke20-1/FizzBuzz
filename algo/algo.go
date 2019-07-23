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

// fizzbuzzAlgo is mainy algo
func FizzbuzzAlgo(fizzbuzz *Fizzbuzz) (resultfizzbuzz string) {
	//resultfizzbuzz = append(resultfizzbuzz, ",")

	for i := 0; i <= fizzbuzz.Limit; i++ {
		if i == 0 {
			resultfizzbuzz += strconv.Itoa(i) + ","
			// resultfizzbuzz += "," + ","
			continue
		}
		if (i%fizzbuzz.Int1 == 0) && (i%fizzbuzz.Int2 == 0) {
			resultfizzbuzz += fizzbuzz.Fizz + fizzbuzz.Buzz + ","
			// resultfizzbuzz += "," + ","
			continue
		}
		if i%fizzbuzz.Int1 == 0 {
			resultfizzbuzz += fizzbuzz.Fizz + ","
			// resultfizzbuzz += "," + ","
			continue
		}
		if i%fizzbuzz.Int2 == 0 {
			resultfizzbuzz += fizzbuzz.Buzz + ","
			// resultfizzbuzz += "," + ","
			continue
		}
		resultfizzbuzz += strconv.Itoa(i) + ","
		// resultfizzbuzz = append(resultfizzbuzz, ",")
	}
	return resultfizzbuzz[:len(resultfizzbuzz)-1]
}

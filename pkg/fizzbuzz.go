package pkg

import (
	"errors"
	"fmt"

	"github.com/tsauzeau/lbc/cmd/lbc/forms"
)

// Fizzbuzz Returns a list of strings with numbers from 1 to limit
// where: all multiples of int1 are replaced by str1
// all multiples of int2 are replaced by str2
// all multiples of int1 and int2 are replaced by str1str2.
func Fizzbuzz(fizzbuzz *forms.FizzbuzzForm) (res []string, err error) {
	if fizzbuzz.Limit <= 0 || fizzbuzz.Int1 <= 0 || fizzbuzz.Int2 <= 0 {
		return nil, errors.New("Limit, Int1 and Int2 needs to be positive values")
	}
	for i := 1; i <= fizzbuzz.Limit+1; i++ {
		if i%(fizzbuzz.Int1*fizzbuzz.Int2) == 0 {
			res = append(res, fmt.Sprintf("%s%s", fizzbuzz.String1, fizzbuzz.String2))
		} else if i%fizzbuzz.Int1 == 0 {
			res = append(res, fmt.Sprintf(fizzbuzz.String1))
		} else if i%fizzbuzz.Int2 == 0 {
			res = append(res, fmt.Sprintf(fizzbuzz.String2))
		} else {
			res = append(res, fmt.Sprintf("%d", i))
		}
	}
	res = append(res, "...")
	return res, nil
}

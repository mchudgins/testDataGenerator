package ssn

import (
	"fmt"
	"math/rand"
)

/*
  This generator uses https://en.wikipedia.org/wiki/Social_Security_number
  as the reference for a SSN.  According to the article,
  "Numbers with all zeros in any digit group (000-##-####, ###-00-####, ###-##-0000)"
  represent invalid SSN's.  This generator only issues invalid SSN's.
*/

func GenerateInvalidSSN(r *rand.Rand) string {
	switch r.Int31n(3) {
	case 0:
		return fmt.Sprintf("000-%02d-%04d", r.Int31n(100), r.Int31n(10000))

	case 1:
		return fmt.Sprintf("%03d-00-%04d", r.Int31n(1000), r.Int31n(10000))

	case 2:
		return fmt.Sprintf("%03d-%02d-0000", r.Int31n(1000), r.Int31n(100))

	default:
		return "000-00-0000"
	}
}

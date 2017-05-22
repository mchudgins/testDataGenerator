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
	switch r.Int31n(9) {
	case 0, 1, 2: /* 6 random digits generated */
		return fmt.Sprintf("000-%02d-%04d", r.Int31n(100), r.Int31n(10000))

	case 3, 4, 5, 6: /* 7 random digits generated */
		return fmt.Sprintf("%03d-00-%04d", r.Int31n(1000), r.Int31n(10000))

	case 7, 8: /* 5 random digits generated */
		return fmt.Sprintf("%03d-%02d-0000", r.Int31n(1000), r.Int31n(100))

	default: /* panic? */
		return "000-00-0000"
	}
}

func GenerateSSN(r *rand.Rand) string {
	return fmt.Sprintf("%03d-%02d-%04d",
		r.Int31n(998) + 1,
		r.Int31n(98) + 1,
		r.Int31n(9998) + 1 )
}

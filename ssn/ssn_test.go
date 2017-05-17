package ssn

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSSNGenerator(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 10; i >= 0; i-- {
		s := GenerateInvalidSSN(r)
		fmt.Printf("SSN: %s\n", s)
	}
}

package command

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func TestNewField(t *testing.T) {

	prime, err := rand.Prime(rand.Reader, 258)
	if err != nil {
		t.Fatal(err.Error())
	}

	field := NewField(prime)
	if field.Prime.Cmp(prime) != 0 {
		t.Error("The field has different prime number:", field)
	}
	if field.Max.Cmp(new(big.Int).Sub(prime, big.NewInt(1))) != 0 {
		t.Error("Max of the field is wrong:", field)
	}

}

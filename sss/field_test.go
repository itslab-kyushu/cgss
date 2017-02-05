package sss

import (
	"crypto/rand"
	"encoding/json"
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

func TestMarshal(t *testing.T) {

	var err error
	field := NewField(big.NewInt(12345))
	data, err := json.Marshal(field)
	if err != nil {
		t.Fatal(err.Error())
	}

	var res Field
	if err = json.Unmarshal(data, &res); err != nil {
		t.Fatal(err.Error())
	}
	if field.Prime.Cmp(res.Prime) != 0 {
		t.Error("Unmarshaled prime is wrong:", res)
	}
	if field.Max.Cmp(res.Max) != 0 {
		t.Error("Unmarshaled max is wrong:", res)
	}

}

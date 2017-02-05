package sss

import (
	"encoding/json"
	"math/big"
)

// Field represents a finite field Z/pZ.
type Field struct {
	// Prime number.
	Prime *big.Int
	// Max is Prime - 1.
	Max *big.Int
}

// compactField defines a field to marshal/unmarshal.
type compactField struct {
	// Prime number.
	Prime *big.Int
}

// NewField creates a new finite field.
func NewField(prime *big.Int) *Field {

	return &Field{
		Prime: prime,
		Max:   new(big.Int).Sub(prime, big.NewInt(1)),
	}

}

// MarshalJSON implements Marshaler interface.
func (f *Field) MarshalJSON() ([]byte, error) {

	aux := compactField{
		Prime: f.Prime,
	}
	return json.Marshal(aux)

}

// UnmarshalJSON implements Unmarshaler interface.
func (f *Field) UnmarshalJSON(data []byte) (err error) {

	var aux compactField
	if err = json.Unmarshal(data, &aux); err != nil {
		return
	}
	f.Prime = aux.Prime
	f.Max = new(big.Int).Sub(aux.Prime, big.NewInt(1))
	return

}

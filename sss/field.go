package sss

import "math/big"

// Field represents a finite field Z/pZ.
type Field struct {
	// Prime number.
	Prime *big.Int
	// Max is Prime - 1.
	Max *big.Int
}

// NewField creates a new finite field.
func NewField(prime *big.Int) *Field {

	return &Field{
		Prime: prime,
		Max:   new(big.Int).Sub(prime, big.NewInt(1)),
	}

}

// TODO: Implementn marshaler/unmarshaler to only save Prime attribute.

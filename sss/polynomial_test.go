package sss

import (
	"math/big"
	"testing"
)

func TestPolynomial(t *testing.T) {

	field := NewField(big.NewInt(37))
	threshold := 2
	polynomial, err := NewPolynomial(field, big.NewInt(0), threshold-1)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(polynomial.Coefficients) != 1 {
		t.Fatal("Dimension of the polynomial is wrong:", polynomial)
	}

	var res *big.Int
	if res = polynomial.Call(big.NewInt(0)); res.Int64() != 0 {
		t.Error("F(0) returns a wrong value:")
	}

	if res = polynomial.Call(big.NewInt(1)); res.Cmp(polynomial.Coefficients[0]) != 0 {
		t.Error("F(1) returns a wrong value:", res)
	}

}

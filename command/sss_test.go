package command

import (
	"math/big"
	"testing"
)

func TestSS1(t *testing.T) {

	var err error

	prime := big.NewInt(37)
	secret := big.NewInt(int64(35))

	size := 10
	threshold := 3

	shares, err := Distribute(prime, secret, size, threshold)
	if err != nil {
		t.Fatal(err.Error())
	}

	res, err := Reconstruct(shares[:threshold])
	if err != nil {
		t.Error(err.Error())
	}
	if res.Cmp(secret) != 0 {
		t.Error("SS1 is broken:", res, secret)
	}

}

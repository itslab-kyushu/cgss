package command

import (
	"fmt"
	"math/big"
)

// Share defines a share of Shamir's Secret Sharing scheme.
type Share struct {
	Key   *big.Int
	Value *big.Int
	Field *Field
}

// Distribute computes shares made by a prime number and a secret.
func Distribute(prime *big.Int, value *big.Int, size, threshold int) (shares []Share, err error) {

	field := NewField(prime)
	polynomial, err := NewPolynomial(field, value, threshold-1)
	if err != nil {
		return
	}

	shares = make([]Share, size)
	for i := range shares {
		key := big.NewInt(int64(i + 1))
		shares[i] = Share{
			Key:   key,
			Value: polynomial.Call(key),
			Field: field,
		}

	}

	return

}

// Reconstruct computes the secret value from a set of shares.
func Reconstruct(shares []Share) (res *big.Int, err error) {

	if len(shares) == 0 {
		err = fmt.Errorf("No shares are given")
		return
	}

	res = big.NewInt(0)
	field := shares[0].Field
	for i, s := range shares {
		res.Add(res, new(big.Int).Mul(s.Value, beta(field, shares, i)))
	}
	res.Mod(res, field.Prime)

	return

}

// beta computes the following value:
//   \mul_{i<=u<=k, u!=t} \frac{u-th key}{(u-th key) - (t-th key)}
func beta(field *Field, shares []Share, t int) *big.Int {

	res := big.NewInt(1)
	for i, s := range shares {
		if i == t {
			continue
		}
		sub := new(big.Int).Mod(new(big.Int).Sub(s.Key, shares[t].Key), field.Prime)
		v := new(big.Int).Mul(s.Key, new(big.Int).ModInverse(sub, field.Prime))
		res.Mul(res, v)
		res.Mod(res, field.Prime)
	}

	return res.Mod(res, field.Prime)

}

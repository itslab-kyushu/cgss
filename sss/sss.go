package sss

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

// Share defines a share of Shamir's Secret Sharing scheme.
type Share struct {
	Key   *big.Int
	Value []*big.Int
	Field *Field
}

// Distribute computes shares having a given secret.
func Distribute(secret []byte, chankByte, size, threshold int) (shares []Share, err error) {

	prime, err := rand.Prime(rand.Reader, chankByte*8+1)
	if err != nil {
		return
	}
	field := NewField(prime)

	nvalue := int(math.Ceil(float64(len(secret)) / float64(chankByte)))
	shares = make([]Share, size)
	for i := range shares {
		key := big.NewInt(int64(i + 1))
		shares[i] = Share{
			Key:   key,
			Value: make([]*big.Int, nvalue),
			Field: field,
		}
	}

	var value *big.Int
	for chank := 0; chank < nvalue; chank++ {
		if len(secret) > chankByte {
			value = new(big.Int).SetBytes(secret[:chankByte])
			secret = secret[chankByte:]
		} else {
			value = new(big.Int).SetBytes(secret)
			secret = nil
		}

		polynomial, err := NewPolynomial(field, value, threshold-1)
		if err != nil {
			return nil, err
		}

		for i := range shares {
			key := big.NewInt(int64(i + 1))
			shares[i].Value[chank] = polynomial.Call(key)
		}

	}

	return
}

// Reconstruct computes the secret value from a set of shares.
func Reconstruct(shares []Share) (bytes []byte, err error) {

	if len(shares) == 0 {
		err = fmt.Errorf("No shares are given")
		return
	}

	bytes = []byte{}
	for chank := 0; chank < len(shares[0].Value); chank++ {

		value := big.NewInt(0)
		field := shares[0].Field
		for i, s := range shares {
			value.Add(value, new(big.Int).Mul(s.Value[chank], beta(field, shares, i)))
		}
		value.Mod(value, field.Prime)

		bytes = append(bytes, value.Bytes()...)

	}
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
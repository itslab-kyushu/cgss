package cgss

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"

	"github.com/itslab-kyushu/cgss/sss"
)

// Share defines a share of the Cross-Group Secret Sharing scheme.
type Share struct {
	GroupShare sss.Share
	DataShare  sss.Share
}

// Distribute computes shares having a given secret.
func Distribute(secret []byte, chunksize int, allocation Allocation, gthreshold, dthreshold int) (shares []Share, err error) {

	// Prepare a field.
	prime, err := rand.Prime(rand.Reader, chunksize*8+2)
	if err != nil {
		return
	}
	field := sss.NewField(prime)

	// Total number of chunks.
	nchunk := int(math.Ceil(float64(len(secret)) / float64(chunksize)))
	nshare := allocation.Sum()

	// Prepare shares.
	shares = make([]Share, nshare)
	for i := range shares {
		shares[i] = Share{
			DataShare: sss.Share{
				Field: field,
				Key:   big.NewInt(int64(i + 1)),
				Value: make([]*big.Int, nchunk),
			},
		}
	}

	var value *big.Int
	for chunk := 0; chunk < nchunk; chunk++ {

		if len(secret) > chunksize {
			value = new(big.Int).SetBytes(secret[:chunksize])
			secret = secret[chunksize:]
		} else {
			value = new(big.Int).SetBytes(secret)
			secret = nil
		}

		// Generate reconstructor's secrets.
		nu, err := rand.Int(rand.Reader, field.Max)
		if err != nil {
			return nil, err
		}

		// Create a tentative secret.
		c := new(big.Int).Add(value, nu)

		// Create shares for the reconstructor's secret.
		rshares, err := sss.Distribute(nu.Bytes(), chunksize+2, allocation.Size(), gthreshold)
		if err != nil {
			return nil, err
		}

		// Create shares for the tentative secret.
		polynomial, err := sss.NewPolynomial(field, c, dthreshold-1)
		if err != nil {
			return nil, err
		}
		iter := allocation.Iterator()
		for i := range shares {
			key := big.NewInt(int64(i + 1))
			shares[i].DataShare.Value[chunk] = polynomial.Call(key)

			group, ok := iter.Next()
			if !ok {
				return nil, fmt.Errorf("Allocation is not enough: %v", allocation)
			}
			if shares[i].GroupShare.Key == nil {
				shares[i].GroupShare.Field = rshares[group].Field
				shares[i].GroupShare.Key = big.NewInt(int64(group) + 1)
				shares[i].GroupShare.Value = make([]*big.Int, nchunk)
			}
			if len(rshares[group].Value) != 1 {
				return nil, fmt.Errorf("Length of group share for a chunk is wrong: %v", rshares)
			}
			shares[i].GroupShare.Value[chunk] = rshares[group].Value[0]

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
	for chunk := 0; chunk < len(shares[0].DataShare.Value); chunk++ {

		value := big.NewInt(0)
		field := shares[0].DataShare.Field
		for i, s := range shares {
			value.Add(value, new(big.Int).Mul(s.DataShare.Value[chunk], beta(field, shares, i)))
		}
		value.Mod(value, field.Prime)

		gshares := distinctGroupShares(shares, chunk)
		nu, err := sss.Reconstruct(gshares)
		if err != nil {
			return nil, err
		}
		value.Sub(value, new(big.Int).SetBytes(nu))

		bytes = append(bytes, value.Bytes()...)

	}
	return

}

// beta computes the following value:
//   \mul_{i<=u<=k, u!=t} \frac{u-th key}{(u-th key) - (t-th key)}
func beta(field *sss.Field, shares []Share, t int) *big.Int {

	res := big.NewInt(1)
	for i, s := range shares {
		if i == t {
			continue
		}
		sub := new(big.Int).Mod(new(big.Int).Sub(s.DataShare.Key, shares[t].DataShare.Key), field.Prime)
		v := new(big.Int).Mul(s.DataShare.Key, new(big.Int).ModInverse(sub, field.Prime))
		res.Mul(res, v)
		res.Mod(res, field.Prime)
	}

	return res.Mod(res, field.Prime)

}

// distinctGroupShares returns a set of distinct group shares.
func distinctGroupShares(shares []Share, index int) (res []sss.Share) {

	set := map[string]sss.Share{}
	for _, s := range shares {
		key := s.GroupShare.Key.Text(16)
		if _, exist := set[key]; !exist {
			set[key] = s.GroupShare
		}
	}

	res = make([]sss.Share, len(set))
	i := 0
	for _, v := range set {
		res[i].Field = v.Field
		res[i].Key = v.Key
		res[i].Value = []*big.Int{v.Value[index]}
		i++
	}

	return

}

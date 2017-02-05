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

	nu := make([]*big.Int, nchunk)
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
		nu[chunk], err = rand.Int(rand.Reader, field.Max)
		if err != nil {
			return
		}

		// Create a tentative secret.
		c := new(big.Int).Add(value, nu[chunk])
		c.Mod(c, field.Prime)

		// Create shares for the reconstructor's secret.
		rshares, err := sss.Distribute(nu[chunk].Bytes(), chunksize+2, allocation.Size(), gthreshold)
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
				shares[i].GroupShare.Key = big.NewInt(int64(group))
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

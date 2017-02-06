//
// cgss/cgss.go
//
// Copyright (c) 2017 Junpei Kawamoto
//
// This file is part of cgss.
//
// cgss is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// cgss is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with cgss.  If not, see <http://www.gnu.org/licenses/>.
//

package cgss

import (
	"context"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"runtime"

	"github.com/itslab-kyushu/cgss/sss"
	"golang.org/x/sync/errgroup"
)

// Share defines a share of the Cross-Group Secret Sharing scheme.
type Share struct {
	GroupShare []sss.Share
	DataShare  sss.Share
}

// GroupKey returns the group key of the share.
func (s *Share) GroupKey() *big.Int {
	if len(s.GroupShare) == 0 {
		return nil
	}
	return s.GroupShare[0].Key
}

// DataKey returns the data key of the share.
func (s *Share) DataKey() *big.Int {
	return s.DataShare.Key
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
			GroupShare: make([]sss.Share, nchunk),
			DataShare: sss.Share{
				Field: field,
				Key:   big.NewInt(int64(i + 1)),
				Value: make([]*big.Int, nchunk),
			},
		}
	}

	var value *big.Int
	wg, ctx := errgroup.WithContext(context.Background())
	cpus := runtime.NumCPU()
	semaphore := make(chan struct{}, cpus)
	for chunk := 0; chunk < nchunk; chunk++ {

		// Check the context.
		select {
		case <-ctx.Done():
			break
		default:
		}

		if len(secret) > chunksize {
			value = new(big.Int).SetBytes(secret[:chunksize])
			secret = secret[chunksize:]
		} else {
			value = new(big.Int).SetBytes(secret)
			secret = nil
		}

		func(chunk int, value *big.Int) {

			semaphore <- struct{}{}
			wg.Go(func() (err error) {
				defer func() { <-semaphore }()

				// Check the context.
				select {
				case <-ctx.Done():
					return
				default:
				}

				// Generate reconstructor's secrets.
				nu, err := rand.Int(rand.Reader, field.Max)
				if err != nil {
					return
				}

				// Create a tentative secret.
				c := new(big.Int).Add(value, nu)

				// Create shares for the reconstructor's secret.
				rshares, err := sss.Distribute(nu.Bytes(), chunksize, allocation.Size(), gthreshold)
				if err != nil {
					return
				}

				// Create shares for the tentative secret.
				polynomial, err := sss.NewPolynomial(field, c, dthreshold-1)
				if err != nil {
					return
				}
				iter := allocation.Iterator()
				for i := range shares {
					key := big.NewInt(int64(i + 1))
					shares[i].DataShare.Value[chunk] = polynomial.Call(key)
					group, ok := iter.Next()
					if !ok {
						return fmt.Errorf("Allocation is not enough: %v", allocation)
					}
					shares[i].GroupShare[chunk] = rshares[group]
				}
				return

			})

		}(chunk, value)

	}

	return shares, wg.Wait()

}

// Reconstruct computes the secret value from a set of shares.
func Reconstruct(shares []Share) (res []byte, err error) {

	if len(shares) == 0 {
		err = fmt.Errorf("No shares are given")
		return
	}

	bytes := make([][]byte, len(shares[0].DataShare.Value))
	wg, ctx := errgroup.WithContext(context.Background())
	semaphore := make(chan struct{}, runtime.NumCPU())
	for chunk := 0; chunk < len(shares[0].DataShare.Value); chunk++ {

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		func(chunk int) {

			semaphore <- struct{}{}
			wg.Go(func() (err error) {
				defer func() { <-semaphore }()

				select {
				case <-ctx.Done():
					return
				default:
				}

				value := big.NewInt(0)
				field := shares[0].DataShare.Field
				for i, s := range shares {
					value.Add(value, new(big.Int).Mul(s.DataShare.Value[chunk], beta(field, shares, i)))
				}
				value.Mod(value, field.Prime)

				gshares, err := distinctGroupShares(shares, chunk)
				if err != nil {
					return
				}
				nu, err := sss.Reconstruct(gshares)
				if err != nil {
					return
				}
				value.Sub(value, new(big.Int).SetBytes(nu))
				value.Mod(value, field.Prime)
				bytes[chunk] = value.Bytes()
				return
			})

		}(chunk)

	}

	if err = wg.Wait(); err != nil {
		return
	}
	for _, v := range bytes {
		res = append(res, v...)
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
		sub := new(big.Int).Mod(new(big.Int).Sub(s.DataKey(), shares[t].DataKey()), field.Prime)
		v := new(big.Int).Mul(s.DataKey(), new(big.Int).ModInverse(sub, field.Prime))
		res.Mul(res, v)
		res.Mod(res, field.Prime)
	}
	return res.Mod(res, field.Prime)
}

// distinctGroupShares returns a set of distinct group shares.
func distinctGroupShares(shares []Share, index int) (res []sss.Share, err error) {
	res = []sss.Share{}
	set := map[string]struct{}{}
	for _, s := range shares {
		key := s.GroupKey()
		if key == nil {
			return nil, fmt.Errorf("Group shares are broken")
		}
		id := key.Text(16)
		if _, exist := set[id]; !exist {
			set[id] = struct{}{}
			res = append(res, s.GroupShare[index])
		}
	}
	return
}

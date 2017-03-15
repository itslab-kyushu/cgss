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
	"io"
	"io/ioutil"
	"math"
	"math/big"
	"runtime"

	pb "gopkg.in/cheggaaa/pb.v1"

	"github.com/itslab-kyushu/sss/sss"
	"golang.org/x/sync/errgroup"
)

// DistributeOpt defines arguments of Distribute function.
type DistributeOpt struct {
	// Secrets will be divided into chunks based on the ChunkSize.
	ChunkSize int
	// Share assignment information.
	Allocation Allocation
	// Minimum number of groups from which shares must be collected to reconstruct secrets.
	GroupThreshold int
	// Minimum number of shares required to reconstruct secrets.
	DataThreshold int
}

// kv defines a simple key-value pair of *big.Int.
type kv struct {
	Key   *big.Int
	Value *big.Int
}

// Distribute computes cross-groupt secret sharing shares of the given secret
// according to the given configuration, opt.
// If status isn't nil, logging information will be written to the Writer.
func Distribute(ctx context.Context, secret []byte, opt *DistributeOpt, status io.Writer) (shares []Share, err error) {

	if status == nil {
		status = ioutil.Discard
	}

	// Prepare a field.
	prime, err := rand.Prime(rand.Reader, opt.ChunkSize*8+2)
	if err != nil {
		return
	}
	field := sss.NewField(prime)

	// Total number of chunks.
	nchunk := int(math.Ceil(float64(len(secret)) / float64(opt.ChunkSize)))
	nshare := opt.Allocation.Sum()

	// Prepare shares.
	shares = make([]Share, nshare)
	for i := range shares {
		shares[i] = Share{
			Field:       field,
			GroupShares: make([]*big.Int, nchunk),
			DataKey:     big.NewInt(int64(i + 1)),
			DataShares:  make([]*big.Int, nchunk),
		}
	}

	// Configure logging.
	bar := pb.New(nchunk)
	bar.Output = status
	bar.Prefix("Chunk")
	bar.Start()
	defer bar.Finish()

	wg, ctx := errgroup.WithContext(ctx)
	cpus := runtime.NumCPU()
	semaphore := make(chan struct{}, cpus)
	for chunk := 0; chunk < nchunk; chunk++ {

		// Check the context.
		select {
		case <-ctx.Done():
			break
		default:
		}

		var value *big.Int
		if len(secret) > opt.ChunkSize {
			value = new(big.Int).SetBytes(secret[:opt.ChunkSize])
			secret = secret[opt.ChunkSize:]
		} else {
			value = new(big.Int).SetBytes(secret)
			secret = nil
		}

		func(chunk int, value *big.Int) {

			semaphore <- struct{}{}
			wg.Go(func() (err error) {
				defer func() { <-semaphore }()
				defer bar.Increment()

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
				rshares, err := distribute(nu, field, opt.Allocation.Size(), opt.GroupThreshold)
				if err != nil {
					return
				}

				// Create shares for the tentative secret.
				polynomial, err := sss.NewPolynomial(field, c, opt.DataThreshold-1)
				if err != nil {
					return
				}
				iter := opt.Allocation.Iterator()
				for i := range shares {
					key := big.NewInt(int64(i + 1))
					shares[i].DataShares[chunk] = polynomial.Call(key)
					group, ok := iter.Next()
					if !ok {
						return fmt.Errorf("Allocation is not enough: %v", opt.Allocation)
					}
					// Set group shares.
					if shares[i].GroupKey == nil {
						shares[i].GroupKey = big.NewInt(int64(group + 1))
					}
					shares[i].GroupShares[chunk] = rshares[group]
				}
				return

			})

		}(chunk, value)

	}

	return shares, wg.Wait()

}

// distribute computes shares for a reconstructor's secret.
func distribute(nu *big.Int, field *sss.Field, size, gthreshold int) (shares []*big.Int, err error) {

	polynomial, err := sss.NewPolynomial(field, nu, gthreshold-1)
	if err != nil {
		return
	}

	shares = make([]*big.Int, size)
	for i := range shares {
		key := big.NewInt(int64(i + 1))
		shares[i] = polynomial.Call(key)
	}
	return

}

// Reconstruct computes the secret from a set of shares.
// The given shares must satisfy the group and data constraints set to
// distribute the secret. If the number of groups and/or the number of shares
// are not enough, the result is undefined (maybe some random value).
//
// If status isn't nil, logging information will be written to the Writer.
func Reconstruct(ctx context.Context, shares []Share, status io.Writer) (res []byte, err error) {

	if status == nil {
		status = ioutil.Discard
	}

	if len(shares) == 0 {
		err = fmt.Errorf("No shares are given")
		return
	}
	nchunk := len(shares[0].DataShares)
	field := shares[0].Field

	// Configure logging.
	bar := pb.New(nchunk)
	bar.Output = status
	bar.Prefix("Chunk")
	bar.Start()
	defer bar.Finish()

	bytes := make([][]byte, nchunk)
	wg, ctx := errgroup.WithContext(ctx)
	semaphore := make(chan struct{}, runtime.NumCPU())
	for chunk := 0; chunk < nchunk; chunk++ {

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		func(chunk int) {

			semaphore <- struct{}{}
			wg.Go(func() (err error) {
				defer func() { <-semaphore }()
				defer bar.Increment()

				select {
				case <-ctx.Done():
					return
				default:
				}

				value := big.NewInt(0)
				for i, s := range shares {
					value.Add(value, new(big.Int).Mul(s.DataShares[chunk], beta(field, shares, i)))
				}
				value.Mod(value, field.Prime)

				gshares, err := distinctGroupShares(shares, chunk)
				if err != nil {
					return
				}
				nu := reconstruct(gshares, field)
				value.Sub(value, nu)
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

// reconstruct computes a reconstructor's secret from a set of shares.
func reconstruct(shares []*kv, field *sss.Field) *big.Int {

	value := big.NewInt(0)
	for i, s := range shares {

		beta := big.NewInt(1)
		for j, t := range shares {
			if i == j {
				continue
			}
			sub := new(big.Int).Mod(new(big.Int).Sub(t.Key, s.Key), field.Prime)
			v := new(big.Int).Mul(t.Key, new(big.Int).ModInverse(sub, field.Prime))
			beta.Mul(beta, v)
			beta.Mod(beta, field.Prime)
		}
		value.Add(value, new(big.Int).Mul(s.Value, beta))

	}

	value.Mod(value, field.Prime)
	return value

}

// beta computes the following value:
//   \mul_{1<=u<=k, u!=t} \frac{u-th key}{(u-th key) - (t-th key)}
func beta(field *sss.Field, shares []Share, t int) *big.Int {
	res := big.NewInt(1)
	for i, s := range shares {
		if i == t {
			continue
		}
		sub := new(big.Int).Mod(new(big.Int).Sub(s.DataKey, shares[t].DataKey), field.Prime)
		v := new(big.Int).Mul(s.DataKey, new(big.Int).ModInverse(sub, field.Prime))
		res.Mul(res, v)
		res.Mod(res, field.Prime)
	}
	return res.Mod(res, field.Prime)
}

// distinctGroupShares returns a set of distinct group shares.
func distinctGroupShares(shares []Share, index int) (res []*kv, err error) {
	res = []*kv{}
	set := map[string]struct{}{}
	for _, s := range shares {
		if s.GroupKey == nil {
			return nil, fmt.Errorf("Group shares are broken")
		}
		id := s.GroupKey.Text(16)
		if _, exist := set[id]; !exist {
			set[id] = struct{}{}
			res = append(res, &kv{
				Key:   s.GroupKey,
				Value: s.GroupShares[index],
			})
		}
	}
	return
}

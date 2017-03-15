//
// cgss/cgss_test.go
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
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestCGSS(t *testing.T) {

	secret := []byte(`abcdefgaerogih:weori:ih:opih:oeijhg@roeinv;dlkjh:
		roihg:3pw9bdlnbmxznd:lah:orsihg:operinbk:sldfj:aporinb`)
	chunksize := 8

	allocation := Allocation{2, 2, 2}
	gthreshold := 2
	dthreshold := 3
	ctx := context.Background()

	shares, err := Distribute(ctx, secret, &DistributeOpt{
		ChunkSize:      chunksize,
		Allocation:     allocation,
		GroupThreshold: gthreshold,
		DataThreshold:  dthreshold,
	}, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(shares) != allocation.Sum() {
		t.Fatal("Number of generated shares is wrong:", len(shares))
	}

	obtained := []Share{shares[0], shares[1], shares[2]}
	res, err := Reconstruct(ctx, obtained, nil)
	if err != nil {
		t.Error(err.Error())
	}
	if string(secret) != string(res) {
		t.Error("Reconstructed secret is wrong")
	}

}

func TestMarshall(t *testing.T) {

	var err error
	secret := []byte("abcdefg")
	chunksize := 8

	allocation := Allocation{1, 1}
	gthreshold := 2
	dthreshold := 2

	shares, err := Distribute(context.Background(), secret, &DistributeOpt{
		ChunkSize:      chunksize,
		Allocation:     allocation,
		GroupThreshold: gthreshold,
		DataThreshold:  dthreshold,
	}, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	bytes, err := json.Marshal(shares[0])
	if err != nil {
		t.Fatal(err.Error())
	}

	var res Share
	if err = json.Unmarshal(bytes, &res); err != nil {
		t.Error(err.Error())
	}

	for i, v := range shares[0].GroupShares {
		if v.Cmp(res.GroupShares[i]) != 0 {
			t.Error("Marshal/Unmarshal don't work as expected:", res)
		}
	}

	for i, v := range shares[0].DataShares {
		if v.Cmp(res.DataShares[i]) != 0 {
			t.Error("Marshal/Unmarshal don't work as expected:", res)
		}
	}

}

// The following example assumes three groups and assigns two shares to
// each of them, and the group and data thredholds are set to 2 and 3,
// respectively. It means, to reconstruct the secret, at least three shares must
// be collected from at least 2 groups.
//
// Since the chunk size is set to 8bytes, the secret will be divided every
// 8bytes and each chunk will be converted to a set of shares.
//
// From the share assignment, the dictribute function makes totally 6 shares,
// and it thus returns a slice of shares of which the length is 6.
// Note that the returned shares do not have any information about groups but
// the order of them are associated with the given allocation.
// More precisely, shares[0] and shares[1] are for the group 1, shares[2] and
// shares[3] are for the groupt 2, and shares[4] and shares[5] are for the
// group 3 in this example.
func ExampleDistribute() {

	secret := []byte(`abcdefgaerogih:weori:ih:opih:oeijhg@roeinv;dlkjh:
		roihg:3pw9bdlnbmxznd:lah:orsihg:operinbk:sldfj:aporinb`)
	chunksize := 8

	allocation := Allocation{2, 2, 2}
	gthreshold := 2
	dthreshold := 3
	ctx := context.Background()

	shares, err := Distribute(ctx, secret, &DistributeOpt{
		ChunkSize:      chunksize,
		Allocation:     allocation,
		GroupThreshold: gthreshold,
		DataThreshold:  dthreshold,
	}, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(len(shares))
	// Output: 6

}

// The following example reconstructs the secret from a subset of distributed
// shares. As same as the example of Distribute function, the secret is
// distributed into 6 shares over 3 groups so that each group has 2 shares.
//
// Since the group threshold and data threshold are 2 and 3, respectively,
// we use two shares from group 1 and one share from group 2, i.e. three shares
// from two groups, to reconstruct the secret.
//
// Reconstruct function returns a byte slice; you might need to cast it to
// string if the secret is a string.
func ExampleReconstruct() {

	ctx := context.Background()
	// Distribute step is as same as the example of Dictribute function.
	secret := []byte("This is secret information")
	shares, err := Distribute(ctx, secret, &DistributeOpt{
		ChunkSize:      8,
		Allocation:     Allocation{2, 2, 2},
		GroupThreshold: 2,
		DataThreshold:  3,
	}, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Pick up two shares from group 1 and one share from group 2.
	subset := []Share{shares[0], shares[1], shares[2]}
	res, err := Reconstruct(ctx, subset, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(res))
	// Output: This is secret information

}

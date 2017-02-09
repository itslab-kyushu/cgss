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

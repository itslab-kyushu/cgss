//
// client/command/remote/conv_test.go
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

package remote

import (
	"context"
	"testing"

	"github.com/itslab-kyushu/cgss/cgss"
)

func TestConvert(t *testing.T) {

	var err error
	secret := []byte("abcdefg")
	chunksize := 8

	allocation := cgss.Allocation{1, 1}
	gthreshold := 2
	dthreshold := 2

	shares, err := cgss.Distribute(context.Background(), secret, &cgss.DistributeOpt{
		ChunkSize:      chunksize,
		Allocation:     allocation,
		GroupThreshold: gthreshold,
		DataThreshold:  dthreshold,
	}, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(shares) != allocation.Sum() {
		t.Fatal("Distribute didn't make enough shares.")
	}

	v := ToValue(&shares[0])
	res, err := FromValue(v)
	if err != nil {
		t.Fatal(err.Error())
	}

	if shares[0].Field.Prime.Cmp(res.Field.Prime) != 0 {
		t.Error("Field is not same:", shares[0].Field, res.Field)
	}
	if shares[0].GroupKey.Cmp(res.GroupKey) != 0 {
		t.Error("GroupKey is not same:", shares[0].GroupKey, res.GroupKey)
	}
	for i, v := range shares[0].GroupShares {
		if v.Cmp(res.GroupShares[i]) != 0 {
			t.Error("GroupShares are not same:", shares[0].GroupShares, res.GroupShares)
		}
	}
	if shares[0].DataKey.Cmp(res.DataKey) != 0 {
		t.Error("DataKey is not same:", shares[0].DataKey, res.DataKey)
	}
	for i, v := range shares[0].DataShares {
		if v.Cmp(res.DataShares[i]) != 0 {
			t.Error("DataShares are not same:", shares[0].DataShares, res.DataShares)
		}
	}

}

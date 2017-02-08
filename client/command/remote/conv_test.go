package remote

import (
	"context"
	"testing"

	"github.com/itslab-kyushu/cgss/cgss"
	"github.com/itslab-kyushu/cgss/sss"
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

	v := ToValue(shares[0])
	res, err := FromValue(v)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !cmpShare(res.DataShare, shares[0].DataShare) {
		t.Error("DataShare is not same:", res.DataShare, shares[0].DataShare)
	}
	for i, v := range res.GroupShare {
		if !cmpShare(v, shares[0].GroupShare[i]) {
			t.Error("GroupShare is not same:", v, shares[0].GroupShare[i])
		}
	}

}

func cmpShare(lhs, rhs sss.Share) bool {

	if lhs.Field.Prime.Cmp(rhs.Field.Prime) != 0 {
		return false
	}
	if lhs.Key.Cmp(rhs.Key) != 0 {
		return false
	}
	if len(lhs.Value) != len(rhs.Value) {
		return false
	}
	for i, v := range lhs.Value {
		if v.Cmp(rhs.Value[i]) != 0 {
			return false
		}
	}
	return true

}

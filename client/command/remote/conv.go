package remote

import (
	"fmt"
	"math/big"

	"github.com/itslab-kyushu/cgss/cgss"
	"github.com/itslab-kyushu/cgss/kvs"
	"github.com/itslab-kyushu/cgss/sss"
)

// ToValue converts a share to a value.
func ToValue(share cgss.Share) *kvs.Value {

	res := &kvs.Value{
		GroupShare: make([]*kvs.Share, len(share.GroupShare)),
		DataShare: &kvs.Share{
			Key:   share.DataShare.Key.Text(16),
			Value: make([]string, len(share.DataShare.Value)),
			Field: &kvs.Field{
				Prime: share.DataShare.Field.Prime.Text(16),
			},
		},
	}

	for i, v := range share.DataShare.Value {
		res.DataShare.Value[i] = v.Text(16)
	}

	for i, g := range share.GroupShare {

		s := &kvs.Share{
			Key:   g.Key.Text(16),
			Value: make([]string, len(g.Value)),
			Field: &kvs.Field{
				Prime: g.Field.Prime.Text(16),
			},
		}
		for j, v := range g.Value {
			s.Value[j] = v.Text(16)
		}
		res.GroupShare[i] = s

	}

	return res

}

// FromValue converts a value to a share.
func FromValue(value *kvs.Value) (*cgss.Share, error) {

	var ok bool
	dkey, ok := new(big.Int).SetString(value.DataShare.Key, 16)
	if !ok {
		return nil, fmt.Errorf("Cannot convert the data share key: %v", value.DataShare.Key)
	}
	dprime, ok := new(big.Int).SetString(value.DataShare.Field.Prime, 16)
	if !ok {
		return nil, fmt.Errorf("Cannot convert the prime of the data share: %v", value.DataShare.Field.Prime)
	}

	res := &cgss.Share{
		GroupShare: make([]sss.Share, len(value.GroupShare)),
		DataShare: sss.Share{
			Key:   dkey,
			Value: make([]*big.Int, len(value.DataShare.Value)),
			Field: sss.NewField(dprime),
		},
	}

	for i, v := range value.DataShare.Value {
		if res.DataShare.Value[i], ok = new(big.Int).SetString(v, 16); !ok {
			return nil, fmt.Errorf("Cannot convert a data share value: %v", v)
		}
	}

	for i, g := range value.GroupShare {

		key, ok := new(big.Int).SetString(g.Key, 16)
		if !ok {
			return nil, fmt.Errorf("Cannot convert a group share key: %v", g.Key)
		}

		prime, ok := new(big.Int).SetString(g.Field.Prime, 16)
		if !ok {
			return nil, fmt.Errorf("Cannot convert a prime number: %v", g.Field.Prime)
		}
		s := sss.Share{
			Key:   key,
			Value: make([]*big.Int, len(g.Value)),
			Field: sss.NewField(prime),
		}

		for j, v := range g.Value {
			if s.Value[j], ok = new(big.Int).SetString(v, 16); !ok {
				return nil, fmt.Errorf("Cannot convert a value: %v", v)
			}
		}
		res.GroupShare[i] = s

	}

	return res, nil

}

//
// client/command/remote/conv.go
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
	"fmt"
	"math/big"

	"github.com/itslab-kyushu/cgss/cgss"
	"github.com/itslab-kyushu/cgss/kvs"
	"github.com/itslab-kyushu/sss/sss"
)

// ToValue converts a share to a value.
func ToValue(share *cgss.Share) (value *kvs.Value) {

	value = &kvs.Value{
		Field:       share.Field.Prime.Text(16),
		GroupKey:    share.GroupKey.Text(16),
		GroupShares: make([]string, len(share.GroupShares)),
		DataKey:     share.DataKey.Text(16),
		DataShares:  make([]string, len(share.DataShares)),
	}
	for i, v := range share.GroupShares {
		value.GroupShares[i] = v.Text(16)
	}
	for i, v := range share.DataShares {
		value.DataShares[i] = v.Text(16)
	}
	return

}

// FromValue converts a value to a share.
func FromValue(value *kvs.Value) (*cgss.Share, error) {

	var ok bool
	field, ok := new(big.Int).SetString(value.Field, 16)
	if !ok {
		return nil, fmt.Errorf("Cannot convert the field: %v", value.Field)
	}
	gkey, ok := new(big.Int).SetString(value.GroupKey, 16)
	if !ok {
		return nil, fmt.Errorf("Cannot convert the group key: %v", value.GroupKey)
	}
	dkey, ok := new(big.Int).SetString(value.DataKey, 16)
	if !ok {
		return nil, fmt.Errorf("Cannot convert the data key: %v", value.DataKey)
	}

	res := &cgss.Share{
		Field:       sss.NewField(field),
		GroupKey:    gkey,
		GroupShares: make([]*big.Int, len(value.GroupShares)),
		DataKey:     dkey,
		DataShares:  make([]*big.Int, len(value.DataShares)),
	}
	for i, v := range value.GroupShares {
		if res.GroupShares[i], ok = new(big.Int).SetString(v, 16); !ok {
			return nil, fmt.Errorf("Cannot convert a group share: %v", v)
		}
	}
	for i, v := range value.DataShares {
		if res.DataShares[i], ok = new(big.Int).SetString(v, 16); !ok {
			return nil, fmt.Errorf("Cannot convert a data share: %v", v)
		}
	}

	return res, nil

}

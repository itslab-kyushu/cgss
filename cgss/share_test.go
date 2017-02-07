//
// cgss/share_test.go
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
	"math/big"
	"testing"

	"github.com/itslab-kyushu/cgss/sss"
)

func TestGroupKey(t *testing.T) {

	var share Share
	var res *big.Int

	share = Share{}
	if res = share.GroupKey(); res != nil {
		t.Error("Empty share returns a wrong group key:", res)
	}

	key := big.NewInt(1234)
	share = Share{
		GroupShare: []sss.Share{
			sss.Share{
				Key: key,
			},
			sss.Share{
				Key: big.NewInt(4567),
			},
		},
	}
	if res = share.GroupKey(); res.Cmp(key) != 0 {
		t.Error("GroupKey returns a wrong group key:", res)
	}

}

func TestDataKey(t *testing.T) {

	key := big.NewInt(2345)
	share := Share{
		DataShare: sss.Share{
			Key: key,
		},
	}
	if res := share.DataKey(); res.Cmp(key) != 0 {
		t.Error("DataKey returns a wrong data key:", res)
	}

}

//
// cgss/share.go
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

	"github.com/itslab-kyushu/sss/sss"
)

// Share defines a share of the Cross-Group Secret Sharing scheme.
type Share struct {
	Field       *sss.Field
	GroupKey    *big.Int
	GroupShares []*big.Int

	// GroupShare []sss.Share
	DataShare sss.Share
}

// // GroupKey returns the group key of the share.
// func (s *Share) GroupKey() *big.Int {
// 	if len(s.GroupShare) == 0 {
// 		return nil
// 	}
// 	return s.GroupShare[0].Key
// }

// DataKey returns the data key of the share.
func (s *Share) DataKey() *big.Int {
	return s.DataShare.Key
}

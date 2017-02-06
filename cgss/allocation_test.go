//
// cgss/allocation_test.go
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

import "testing"

func TestAllocationSum(t *testing.T) {

	a := Allocation{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if res := a.Sum(); res != 55 {
		t.Error("sum returns a wrong value:", res)
	}

}

func TestIterator(t *testing.T) {

	a := Allocation{1, 2, 3, 4, 5}
	iter := a.Iterator()

	var (
		v  int
		ok bool
	)
	counter := map[int]int{}
	for {
		v, ok = iter.Next()
		if !ok {
			break
		}
		if _, ok := counter[v]; !ok {
			counter[v] = 0
		}
		counter[v]++
	}

	for k, v := range counter {
		if k+1 != v {
			t.Error("Iterator doesn't work:", counter)
		}
	}

}

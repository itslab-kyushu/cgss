//
// cgss/allocation.go
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

// Allocation defines a share allocation.
type Allocation []int

// Sum returns the summation of this allocate values.
func (a Allocation) Sum() (res int) {
	for _, v := range a {
		res += v
	}
	return
}

// Size returns the size of allocated regions.
func (a Allocation) Size() int {
	return len(a)
}

// Iterator returns an iterator of the assignment.
func (a Allocation) Iterator() *Iterator {
	return newIterator(a)
}

// Iterator defines a set of private members for an iterator of an allocation.
type Iterator struct {
	allocation Allocation
	group      int
	count      int
}

func newIterator(a Allocation) *Iterator {
	return &Iterator{
		allocation: a,
	}
}

// Next returns a next group number.
func (i *Iterator) Next() (int, bool) {

	if i.group >= len(i.allocation) {
		return 0, false
	}

	res := i.group
	i.count++
	if i.count == i.allocation[i.group] {
		i.count = 0
		i.group++
	}
	return res, true

}

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

// Size retunrs the size of allocated regions.
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

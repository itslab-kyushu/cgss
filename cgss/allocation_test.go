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

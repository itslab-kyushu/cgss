package cgss

import "testing"

func TestDistribution(t *testing.T) {

	secret := []byte("abcdefg")
	chunksize := 8

	allocation := Allocation{2, 2, 2}
	gthreshold := 2
	dthreshold := 3

	_, err := Distribute(secret, chunksize, allocation, gthreshold, dthreshold)
	if err != nil {
		t.Fatal(err.Error())
	}

}

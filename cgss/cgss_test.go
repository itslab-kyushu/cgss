package cgss

import "testing"

func TestCGSS(t *testing.T) {

	secret := []byte("abcdefg")
	chunksize := 8

	allocation := Allocation{2, 2, 2}
	gthreshold := 2
	dthreshold := 3

	shares, err := Distribute(secret, chunksize, allocation, gthreshold, dthreshold)
	if err != nil {
		t.Fatal(err.Error())
	}

	obtained := []Share{shares[0], shares[1], shares[2]}
	res, err := Reconstruct(obtained)
	if err != nil {
		t.Error(err.Error())
	}
	if string(secret) != string(res) {
		t.Error("Reconstructed secret is wrong:", string(res))
	}

}

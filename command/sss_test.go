package command

import "testing"

func TestSS1(t *testing.T) {

	var err error

	secret := "abcdefghijklmnopqrstuvwxyz"

	size := 10
	threshold := 3

	shares, err := Distribute([]byte(secret), 8, size, threshold)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(shares)

	res, err := Reconstruct(shares[:threshold])
	if err != nil {
		t.Error(err.Error())
	}
	if string(res) != secret {
		t.Error("SS1 is broken:", res, secret)
	}

}

func TestSS1Word(t *testing.T) {

	var err error

	secret := "a"

	size := 10
	threshold := 3

	shares, err := Distribute([]byte(secret), 256, size, threshold)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(shares)

	res, err := Reconstruct(shares[:threshold])
	if err != nil {
		t.Error(err.Error())
	}
	if string(res) != secret {
		t.Error("SS1 is broken:", res, secret)
	}

}

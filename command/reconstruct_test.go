package command

import "testing"

func TestOutputName(t *testing.T) {

	var res string
	if res = outputFile("simple.2.F.json"); res != "simple" {
		t.Error("Returned filename is wrong:", res)
	}

	if res = outputFile("simple.dat.3.B.json"); res != "simple.dat" {
		t.Error("Returned filename is wrong:", res)
	}

	if res = outputFile(".hidden.2.1.json"); res != ".hidden" {
		t.Error("Returned filename is wrong:", res)
	}

	if res = outputFile("./complex/case.3.4.json"); res != "./complex/case" {
		t.Error("Returned filename is wrong:", res)
	}

	if res = outputFile("dir/complex/case.dat.D.F.json"); res != "dir/complex/case.dat" {
		t.Error("Returned filename is wrong:", res)
	}

	if res = outputFile("invalid.json"); res != "" {
		t.Error("Returned filename is wrong:", res)
	}

}

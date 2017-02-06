//
// cgss/cgss_test.go
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

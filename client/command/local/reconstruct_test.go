//
// client/command/local/reconstruct_test.go
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

package local

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

//
// sss/field.go
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

package sss

import (
	"encoding/json"
	"math/big"
)

// Field represents a finite field Z/pZ.
type Field struct {
	// Prime number.
	Prime *big.Int
	// Max is Prime - 1.
	Max *big.Int
}

// compactField defines a field to marshal/unmarshal.
type compactField struct {
	// Prime number.
	Prime *big.Int
}

// NewField creates a new finite field.
func NewField(prime *big.Int) *Field {

	return &Field{
		Prime: prime,
		Max:   new(big.Int).Sub(prime, big.NewInt(1)),
	}

}

// MarshalJSON implements Marshaler interface.
func (f *Field) MarshalJSON() ([]byte, error) {

	aux := compactField{
		Prime: f.Prime,
	}
	return json.Marshal(aux)

}

// UnmarshalJSON implements Unmarshaler interface.
func (f *Field) UnmarshalJSON(data []byte) (err error) {

	var aux compactField
	if err = json.Unmarshal(data, &aux); err != nil {
		return
	}
	f.Prime = aux.Prime
	f.Max = new(big.Int).Sub(aux.Prime, big.NewInt(1))
	return

}

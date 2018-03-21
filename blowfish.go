/*
	blowfish.go:  Blowfish algorithm implementation in Go.

	Based on C implementation of Paul Kocher.

	Copyright (C) 2018 by Piotr Pszczółkowski (piotr@beesoft.pl)

	This library is free software; you can redistribute it and/or
	modify it under the terms of the GNU Lesser General Public
	License as published by the Free Software Foundation; either
	version 2.1 of the License, or (at your option) any later version.
	This library is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
	Lesser General Public License for more details.
	You should have received a copy of the GNU Lesser General Public
	License along with this library; if not, write to the Free Software
	Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA  02111-1307  USA

	If you require this code under a license other than LGPL, please ask.
*/

package blowfish
 
const (
	N = 16
)

type Blowfish struct {
	P [16 + 2]uint32
	S [4][256]uint32
}

//
// New
//
func New(key []byte) *Blowfish {
	bf := &Blowfish{}

	keyLen := len(key)

	for i := 0; i < 4; i++ {
		for j := 0; j < 256; j++ {
			bf.S[i][j] = ORIG_S[i][j]
		}
	}

	k := 0
	for i := 0; i < (N + 2); i++ {
		data := uint32(0)
		for j := 0; j < 4; j++ {
			data = (data << 8) | uint32(key[k])
			k += 1
			if k >= keyLen {
				k = 0
			}
		}
		bf.P[i] = ORIG_P[i] ^ data
	}

	datal := uint32(0)
	datar := uint32(0)

	for i := 0; i < (N + 2); i += 2 {
		bf.Encrypt(&datal, &datar)
		bf.P[i] = datal
		bf.P[i + 1] = datar
	}

	for i := 0; i < 4; i++ {
		for j:= 0; j < 256; j += 2 {
			bf.Encrypt(&datal, &datar)
			bf.S[i][j] = datal
			bf.S[i][j + 1] = datar
		}
	}
	return bf
}

//
// Encrypt
//
func (bf *Blowfish) Encrypt(xl, xr *uint32) {
	Xl := *xl
	Xr := *xr

	for i := 0; i < N; i++ {
		Xl = Xl ^ bf.P[i]
		Xr = bf.f(Xl) ^ Xr
		Xl, Xr = Xr, Xl
	}

	Xl, Xr = Xr, Xl
	Xr = Xr ^ bf.P[N]
	Xl = Xl ^ bf.P[N + 1]
	*xl = Xl
	*xr = Xr
}

//
// Decrypt
//
func (bf *Blowfish) Decrypt(xl, xr *uint32) {
	Xl := *xl
	Xr := *xr

	for i := (N + 1); i > 1; i-- {
		Xl = Xl ^ bf.P[i]
		Xr = bf.f(Xl) ^ Xr
		Xl, Xr = Xr, Xl
	}
	
	Xl, Xr = Xr, Xl
	Xr = Xr ^ bf.P[1]
	Xl = Xl ^ bf.P[0]
	*xl = Xl
	*xr = Xr
}

/********************************************************************
*                                                                   *
*                           P R I V A T E                           *
*                                                                   *
********************************************************************/

//
// f - helper function
//
func (bf *Blowfish) f(x uint32) uint32 {
	d := uint16(x & 0xff); 	x >>= 8
	c := uint16(x & 0xff); 	x >>= 8
	b := uint16(x & 0xff); 	x >>= 8
	a := uint16(x & 0xff)

	y := bf.S[0][a] + bf.S[1][b]
	y = y ^ bf.S[2][c]
	y = y + bf.S[3][d]
	return y
}

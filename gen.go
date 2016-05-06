/*
    Copyright (C) 2016 cacahuatl < cacahuatl at autistici dot org >

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"math"
	"strings"
	"crypto/rand"
	"math/big"
)

type PassGen struct {
	wordlist []string
}

func NewPassGen() *PassGen {
	return &PassGen{
		wordlist: chbs,
	}
}

func (p *PassGen) Generate(w int) string {
	l := len(p.wordlist)
	var c []string
	for i := 0; i < w; i++ {
		r := rnd(l)
		c = append(c, p.wordlist[r])
	}
	return strings.Join(c, " ")
}

//calculates the shannon entropy of the password, in bits
//this assumes that the attacker knows everything about the password generation
//except for the random choices made from the wordlist in their attack
func (p *PassGen) EntropyEstimate(w int) float64 {
	wc := float64(len(p.wordlist))
	wn := float64(w)
	return math.Log2(math.Pow(wc, wn))
}

func rnd(m int) int {
	bm := big.NewInt(int64(m))
	s, e := rand.Int(rand.Reader, bm)
	if e != nil {
		panic("no random")
	}
	return int(s.Int64())
}

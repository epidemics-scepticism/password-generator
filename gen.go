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

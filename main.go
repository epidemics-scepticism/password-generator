package main

import (
	"fmt"
	"flag"
)

func main() {
	w := flag.Int("count", 5, "Number of words per passphrase")
	d := flag.Bool("dry", false, "Only print the estimated entropy")
	flag.Parse()
	p := NewPassGen()
	fmt.Println("Estimated Entropy", p.EntropyEstimate(*w), "bits")
	if *d == false {
		fmt.Println(p.Generate(*w))
	}
}

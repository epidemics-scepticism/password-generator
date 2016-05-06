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

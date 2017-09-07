package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	fmt.Println(ToBase10("20", 3))
	fmt.Println(FromBase10(27, 16))
	fmt.Println(convertBase("12012", 3, 10))
}

func convertBase(s string, b1 int, b2 int) (r string) {
	a := ToBase10(s, b1)
	r = FromBase10(a, b2)
	return r
}

func ToBase10(s string, b1 int) (d int) {
	var result float64

	for i, r := range s {
		intr := 0
		if r <= '9' {
			intr = int(r - '0')

		} else {
			intr = int(r-'A') * 10
		}
		if intr >= b1 {
			log.Fatal("Invalid digit: ",string(r), " is not part of base ",b1)
		}
		result += math.Pow(float64(b1), float64(len(s)-1-i)) * float64(intr)
	}

	d = int(result)

	return int(d)
}

func FromBase10(d int, b2 int) (s string) {
	for {
		if d == 0 {
			break
		}
		intr := d%b2
		var r rune
		if intr <= 9 {
			r = rune(intr + '0')
		} else {
			r = rune((intr-10)+'A')
		}
		s = string(r) + s
		d = d / b2

	}
	return s
}

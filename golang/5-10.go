package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Reverse(-27))
	fmt.Println(Reverse(314))
}

func Reverse(d int) (r int) {
	factors := make([]int, 0, 0)
	isNeg := false
	if d < 0 {
		isNeg = true
		d *= -1
	}
	for {
		if d == 0 {
			break
		} else {
			factors = append(factors, d%10)
			d = d / 10
		}
	}

	for i, v := range factors {
		r += v * int(math.Pow10(len(factors)-1-i))
	}
	if isNeg {
		r *= -1
	}
	return r
}

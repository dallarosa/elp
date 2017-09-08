package main

import (
	"fmt"
)

func main() {
	fmt.Println(power_iterative(2, 10))
	fmt.Println(power(2, 10))
}

//recursive version
func power(x float32, y int) float32 {
	var a float32
	a = 1
	if y/2 < 4 {
		b := float32(1)
		for i := y; i > 0; i-- {
			b *= x
		}
		return b
	} else {
		a = power(x, y/2) * power(x, y/2)
		if float32(y%2) > 0 {
			a *= x
		}
	}
	return a
}

func power_iterative(x float32, y int) float32 {
	result := float32(1)
	power := y
	a := x
	for ; power > 0; power >>= 1 {
		if power % 2 == 1 {
		result *= a
		}
		a *= a
	}
	return result
}

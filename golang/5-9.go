package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(ColumnToNumber("ABC"))
}

func ColumnToNumber(s string) (d int) {
	for i, r := range s {
		intr := int(r - 'A' + 1)
		d += intr * int(math.Pow(float64(26), float64(len(s)-1-i)))
	}
	return d
}

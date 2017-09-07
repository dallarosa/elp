package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(ToBase10("12012", 3))
	fmt.Println(FromBase10(10, 2))
	fmt.Println(convertBase("12012",3,10))
}

func convertBase(s string, b1 int, b2 int) (r string) {
	a := ToBase10(s,b1)
	r = FromBase10(a,b2)
	return r
}

func ToBase10(s string, b1 int) (d int) {
	numMap := make(map[string]int, 16)
	numMap["0"] = 0
	numMap["1"] = 1
	numMap["2"] = 2
	numMap["3"] = 3
	numMap["4"] = 4
	numMap["5"] = 5
	numMap["6"] = 6
	numMap["7"] = 7
	numMap["8"] = 8
	numMap["9"] = 9
	numMap["A"] = 10
	numMap["B"] = 11
	numMap["C"] = 12
	numMap["D"] = 13
	numMap["E"] = 14
	numMap["F"] = 15

	var result float64

	for i, r := range s {
		result += math.Pow(float64(b1), float64(len(s)-1-i)) * float64(numMap[string(r)])
	}

	d = int(result)

	return int(d)
}

func FromBase10(d int, b2 int) (s string) {
	numMap := make(map[int]string, 16)
	numMap[0] = "0"
	numMap[1] = "1"
	numMap[2] = "2"
	numMap[3] = "3"
	numMap[4] = "4"
	numMap[5] = "5"
	numMap[6] = "6"
	numMap[7] = "7"
	numMap[8] = "8"
	numMap[9] = "9"
	numMap[10] = "A"
	numMap[11] = "B"
	numMap[12] = "C"
	numMap[13] = "D"
	numMap[14] = "E"
	numMap[15] = "F"
	for {
		if d == 0 {
			break
		}
		s = strconv.Itoa(d % b2) + s
		d = d / b2

	}
	return s
}

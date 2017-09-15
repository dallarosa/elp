package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	A := []int{4,3,7,1,10,9,5,6,2,8,7}
	DutchPartition(A,7)
	fmt.Println(A)
}

func DutchPartition(A []int, pivot int) {
	var smaller, equal, bigger int
	bigger = len(A) - 1

	for {
		if equal > bigger {
			break
		} else {
			if A[equal] < pivot {
				swap(A, equal, smaller)
				equal++
				smaller++
			} else if A[equal] == pivot {
				equal++
			} else {
				swap(A, equal, bigger)
				bigger--
			}
		}
	}
}

func swap(A []int, a int, b int) {
	tmp := A[a]
	A[a] = A[b]
	A[b] = tmp
}

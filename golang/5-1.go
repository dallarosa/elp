package main

func getBitSequence(input uint, pos uint) (seq uint) {
	shift := uint(4 * pos)
	mask := uint(0xf)
	seq = (input & (mask << shift)) >> shift
	return seq
}

func getParity(input int) (parity int) {
	bitmap := make(map[int]int)
	bitmap[0] = 0
	bitmap[1] = 1
	bitmap[2] = 1
	bitmap[3] = 2
	bitmap[4] = 1
	bitmap[5] = 2
	bitmap[6] = 2
	bitmap[7] = 3
	bitmap[8] = 1
	bitmap[9] = 2
	bitmap[10] = 2
	bitmap[11] = 3
	bitmap[12] = 2
	bitmap[13] = 3
	bitmap[14] = 3
	bitmap[15] = 4

	a := 0

	for {
		if input > 0 {
			a = a + bitmap[input&0xF]
			input = input >> 4
		} else {
			break
		}
	}
	return a % 2
}

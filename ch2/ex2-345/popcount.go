package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Exercise 2.3
// Rewrite PopCount to use a loop instead of a single expression.
func PopCountByLoop(x uint64) int {
	var count byte
	for i := uint(0); i < 8; i++ {
		count += pc[byte(x>>(i*8))]
	}
	return int(count)
}

// Exercise 2.4
// Write a version of PopCount that counts bits by shifting its argument
//+ through 64 bit positions, testing the rightmost bit each time.
func PopCountByShift(x uint64) int {
	var count int
	for x != 0 {
		if x&1 == 1 {
			count++
		}
		x = x >> 1
	}
	return count
}

// Exercise 2.5
// The x&(x-1) clears the rightmost non-zero bit of x. Write a version
//+ of PopCount that counts bits by using this fact.
func PopCountByClearing(x uint64) int {
	var count int
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

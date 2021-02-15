package main

import "fmt"

// getBit retrieves the bit value from an integer at bit i.
func getBit(num int, i int) bool {
	return (num & (1 << i)) != 0
}

// setBit flips the bit value at bit i.
func setBit(num int, i int) int {
	return num | (1 << i)
}

// clearBit sets the bit at i to 0.
func clearBit(num int, i int) int {
	// create a mask where only the value at i is 0 and everything else is 1
	// 100..0 -> 011..1
	mask := ^1 << i
	// perform an and on our number with the mask
	return num & mask
}

// clearBitsMSBThroughI sets all bit values from the MSB to i (inclusive) to 0
func clearBitsMSBThroughI(num int, i int) int {
	// create a mask where the bits from the MSB through i are 0 and everything past i is 1
	mask := (1 << i) - 1
	return num & mask
}

// clearBitsIThroughLSB sets all bit values from i to the LSB (inclusive) to 0
func clearBitsIThroughLSB(num int, i int) int {
	// create a mask where the bits from i to the LSB are 0 and everything before i is 1
	mask := -1 << (i + 1)
	return num & mask
}

// updateBit sets the value at bit i to either a 1 or a 0 based on the third parameter.
func updateBit(num int, i int, valIs1 bool) int {
	// set our bit value
	val := 0
	if valIs1 {
		val = 1
	}
	// create a mask to clear the ith bit in our number
	clearMask := ^(1 << i)
	// create a mask to set the ith bit in our number
	valueMask := val << i
	return (num & clearMask) | valueMask
}

func main() {
	fmt.Printf("%v\n", updateBit(15, 3, false))
}

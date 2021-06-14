package main

import "fmt"

// [12, 5, 10], n = 3
// iter 1: currIndex = 0, minIndex = 0; lookAhead = 1
// 		1 < 3 <TRUE>
//			5 < 12 <TRUE>
//				minIndex = 1
//			lookAhead = 2
//		2 < 3 <TRUE>
//			10 < 5 <FALSE>
//			lookAhead = 3
// 		3 < 3 <FALSE>
//		[5, 12, 10]
// ..
// ..
// [5, 10, 12]

// Iterates over an array while putting the smallest values into their correct place.
// Time complexity: O(n^2)
// Space complexity: O(1)
// Uses: The good thing about selection sort is it never makes more than O(n)
// swaps and can be useful when memory write is a costly operation.
func selectionSort(arr []int) {
	n := len(arr)

	// iterate over the array
	for currIndex := 0; currIndex < n-1; currIndex++ {
		// index of the minimum item
		minIndex := currIndex
		// look ahead
		lookAhead := currIndex + 1

		// loop while our look ahead is within bounds
		for lookAhead < n {
			// check if our look ahead value is less than our minimum index value
			if arr[lookAhead] < arr[minIndex] {
				// update minimum index
				minIndex = lookAhead
			}
			// increment our look ahead
			lookAhead++
		}

		// swap our minimum index value and our current index value
		arr[minIndex], arr[currIndex] = arr[currIndex], arr[minIndex]
	}
}

func main() {
	arr := []int{12, 5, 10}

	selectionSort(arr)

	fmt.Println(arr)
}

package main

import "fmt"

// [12, 11, 13, 5, 6]
// iter 1: i = 1, currElem = 11, j = 0
// 		0 >= 0 && {12} > {11} <TRUE>
//		[12, 12, 13, 5, 6]
//		j = -1
// [11, 12, 13, 5, 6]
// iter 2: i = 2, currElem = 13, j = 1
//		1 >= 0 && {12} > {13} <FALSE>
// iter 3: i = 3, currElem = 5, j = 2
// 		2 >= 0 && {13} > {5} <TRUE>
//		[11, 12, 13, 13, 6]
//		j = 1
// 		1 >= 0 && {12} > 5 <TRUE>
//		[11, 12, 12, 13, 6]
//		j = 0
// 		0 >= 0 && {11} > {5} <TRUE>
// 		[11, 11, 12, 13, 6]
//		-1 >= 0 && ... <FALSE>
// [5, 11, 12, 13, 6]
// iter 4: i = 4, currElem = 6, j = 3
//		4 >= 0 && {13} > {6} <TRUE>
// 		[5, 11, 12, 13, 13]
//..
//..
//		[5, 11, 11, 12, 13]
// 		0 >= 0 && {5} > {6} <FALSE>
//		[5, 6, 11, 12, 13]
// DONE!

// Iterates over an array and shifts smaller elements to the left.
// Best/Average/Worse runtimes: O(n), O(n^2)
// Space complexity: O(1)
// Uses: Insertion sort is used when number of elements is small. It can also be useful when input array is almost
// sorted, only few elements are misplaced in complete big array.
func insertionSort(arr []int) {
	// loop over the whole array from 1 to n
	for i := 1; i < len(arr); i++ {
		// current element & index of the last element
		currElem := arr[i]
		j := i - 1

		// while j is within the lower bonds & an element behind the current element is lesser
		for j >= 0 && arr[j] > currElem {
			// move the value at j up one
			arr[j+1] = arr[j]
			// decrement j
			j--
		}
		// set the value right after j to the current element
		arr[j+1] = currElem
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6}

	insertionSort(arr)

	fmt.Println(arr)
}

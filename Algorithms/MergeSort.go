package main

import "fmt"

// MergeSort sorts elements in an array in ascending order.
func MergeSort(arr []int, left, right int) {
	if right > left {
		// get the middle index
		middle := (left + (right - 1)) / 2
		// call merge sort with the left half
		MergeSort(arr, left, middle)
		// call merge sort with the right half
		MergeSort(arr, middle+1, right)
		// combine the two halves
		merge(arr, left, middle, right)
	}
}

// Merge combines two sub arrays into sorted order. Assumes arr[l..m] and arr[m+1..r] are sorted.
// Best/Average/Worse runtimes: O(nlogn)
// Space complexity: O(n)
func merge(arr []int, left, middle, right int) {
	// number of elements in the left and right sub arrays
	numLeft, numRight := middle+1-left, right-middle
	leftSubArr, rightSubArr := make([]int, numLeft), make([]int, numRight)
	// indexes for our sub arrays and main array
	leftIdx, rightIdx, arrIdx := 0, 0, left

	// populate left sub array
	for i := 0; i < numLeft; i++ {
		leftSubArr[i] = arr[i+left]
	}
	// populate right sub array
	for i := 0; i < numRight; i++ {
		rightSubArr[i] = arr[i+middle+1]
	}

	// loop while we still have elements from our left & right sub arrays
	for leftIdx < numLeft && rightIdx < numRight {
		// determine which element from our sub arrays should be placed
		if leftSubArr[leftIdx] <= rightSubArr[rightIdx] {
			arr[arrIdx] = leftSubArr[leftIdx]
			leftIdx++
		} else {
			arr[arrIdx] = rightSubArr[rightIdx]
			rightIdx++
		}
		arrIdx++
	}

	// insert remaining elements from our left sub array
	for leftIdx < numLeft {
		arr[arrIdx] = leftSubArr[leftIdx]
		leftIdx++
		arrIdx++
	}

	// insert remaining elements from our right sub array
	for rightIdx < numLeft {
		arr[arrIdx] = rightSubArr[rightIdx]
		rightIdx++
		arrIdx++
	}
}

func main() {
	testArr := []int{4, 3, 5, 1, 8, 5}

	MergeSort(testArr, 0, len(testArr)-1)

	fmt.Printf("%v", testArr)
}

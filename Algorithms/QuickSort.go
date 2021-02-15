package main

import "fmt"

// Partition takes an array and sorts it to have it's partition element (arr[high]) at it's sorted index and all
// less value elements to the left and all greater elements to the right
func partition(arr []int, low int, high int) int {
	// get right most element in the array
	pivot := arr[high]
	// index where our pivot (in the sorted array) lands
	pivotIndex := low - 1

	// loop through the array
	for j := low; j <= high; j++ {
		// current element is smaller than the pivot
		if arr[j] < pivot {
			// increment pivot index
			pivotIndex++
			// move the pivot index element (bigger) to the right
			// and the current element (smaller) to the left
			arr[pivotIndex], arr[j] = arr[j], arr[pivotIndex]
		}
	}
	// swap element at the pivot index with the pivot element
	arr[pivotIndex+1], arr[high] = arr[high], arr[pivotIndex+1]

	// return the pivot index
	return pivotIndex + 1
}

// Quicksort sorts elements in an array in ascending order.
func Quicksort(arr []int, low int, high int) {
	if low < high {
		// sort the pivot element (arr[high]) to it's correct location
		pi := partition(arr, low, high)
		// run quicksort on the two halves
		Quicksort(arr, low, pi-1)
		Quicksort(arr, pi+1, high)
	}
}

func main() {
	testArr := []int{4, 3, 5, 1, 8, 5}

	Quicksort(testArr, 0, len(testArr)-1)

	fmt.Printf("%v", testArr)
}

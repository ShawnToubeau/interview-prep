package main

import "fmt"

func heapify(arr []int, n int, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]

		heapify(arr, n, largest)
	}
}

func heapsort(arr []int) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	fmt.Println(arr)

}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}

	heapsort(arr)
}

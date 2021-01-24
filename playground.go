package main

import "fmt"

func partition(arr []int, low int, high int) int {
	pivot := arr[high]

	i := low - 1

	for j := low; j <= high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}

func quicksort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)

		quicksort(arr, low, pi-1)
		quicksort(arr, pi+1, high)
	}
}

// want nums1 to be smaller than nums2
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	set := make(map[int]int)
	list := make([]int, 0)

	for _, elem := range nums1 {
		set[elem] = set[elem] + 1
	}

	fmt.Printf("Num1 Freq: %v\n", set)

	for _, elem := range nums2 {
		count := set[elem]

		if count > 0 {
			list = append(list, elem)
			set[elem]--
		}
	}

	fmt.Printf("List: %v\n", list)

	return list
}

/**
[1,2,2,1]
[2,2] 0 -> 2; [no, yes]; offset: 0; matches: 1

[1,2,2,1]
  [2,2] 0 -> 2; [yes, yes]; offset: 1; matches: 2

[1,2,2,1]
    [2,2] 0 -> 2; [yes, no]: offset: 2; matches: 1

*/

func main() {
	//nums1 := []int{1, 2, 2, 1}
	//nums2 := []int{2, 2}

	nums1 := []int{4,9,5,9}
	nums2 := []int{9,4,9,8,4}

	fmt.Printf("Res %v\n", intersect(nums1, nums2))
}

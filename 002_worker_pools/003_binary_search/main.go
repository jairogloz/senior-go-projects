package main

import "fmt"

// Begin with low at the begining of the array
// Iterate while low <= high
// For each iteration:
//  - find the middle
//  - if arr[middle] == target: return target
//  - if arr[middle] > target: high = mid // search the left half
//  - if arr[middle] < target: low = mid // search the right half
// If not found return -1

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(binarySearch(a, 1))
	fmt.Println(binarySearch(a, 4))
	fmt.Println(binarySearch(a, 9))

}

func binarySearch(a []int, target int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := (low + high) / 2

		if a[mid] == target {
			return mid
		} else if a[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

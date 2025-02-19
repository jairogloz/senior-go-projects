package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("7: ", binarySearch(a, 7))
	fmt.Println("2: ", binarySearch(a, 2))
	fmt.Println("10: ", binarySearch(a, 10))

}

func binarySearch(a []int, b int) (elementIndex int) {
	minIndex := 0
	maxIndex := len(a) - 1

	for minIndex < maxIndex {
		// Finds the mid spot
		mid := (minIndex + maxIndex) / 2
		// Compares the looked for value with the value at mid
		if a[mid] == b {
			// If equal it returns: we've found the searched element
			return mid
		}
		// If searched < a[mid] : adjust max
		if b < a[mid] {
			maxIndex = mid - 1
		}
		// If searched > a[mid] : adjust min
		if b > a[mid] {
			minIndex = mid + 1
		}
	}

	return -1
}

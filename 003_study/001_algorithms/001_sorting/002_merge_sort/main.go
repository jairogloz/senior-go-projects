package main

import "fmt"

func main() {
	//a := []int{9, 4, 54, 1, 3, 2, 65, 34, 75, 63, 23, 34, 6, -100, -1}
	a := []int{1, 2, 3, 4}
	a2 := mergeSort(a)
	fmt.Println(a)
	fmt.Println(a2)
}

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	// Find the half point
	halfPoint := len(a) / 2
	// Apply mergeSort to left half
	left := mergeSort(a[:halfPoint])
	// Apply mergeSort to right half
	right := mergeSort(a[halfPoint:])
	// Merge
	return merge(left, right)
}

func merge(left, right []int) []int {
	// Iterate over each element of both arrays
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	return append(result, append(left[i:], right[j:]...)...)
}

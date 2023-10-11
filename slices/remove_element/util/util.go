package util

import "slices"

// RemoveLoop
// Using a loop to create a new slice.
func RemoveLoop(slice []int, index int) []int {
	result := make([]int, 0, len(slice)-1)
	for i := range slice {
		if i == index {
			continue
		}
		result = append(result, slice[i])
	}
	return result
}

// RemoveNewSlice
// Creating a new slice and copying element
func RemoveNewSlice(slice []int, index int) []int {
	result := make([]int, len(slice)-1)
	copy(result[:index], slice[:index])
	copy(result[index:], slice[index+1:])
	return result
}

// RemoveCopy removes an element from a slice using copy and slicing.
func RemoveCopy(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

// RemoveAppend removes an element from a slice using append and slicing.
func RemoveAppend(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// RemoveWithSlicesPackage removes an element from a slice using the slices package.
func RemoveWithSlicesPackage(slice []int, index int) []int {
	return slices.Delete(slice, index, index+1)
}

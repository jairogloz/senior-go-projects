package main

import "fmt"

func main() {
	a := []int{3, 6, 1, 2, 9, 3, 5, 1, 3, 5, 7, 12, -7, 84, 98, 10, 15}

	fmt.Println("Before sorting: ", a)

	sortedArray := quickSort(a)

	fmt.Println("After sorting: ", sortedArray)

}

// quickSort
func quickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	// Choose pivot
	pivot := a[len(a)/2]
	// Partition into three sub-slices
	var left, right, equalToPivot []int
	for _, element := range a {
		if element < pivot {
			// - less than pivot
			left = append(left, element)
		} else if element > pivot {
			// - larger than pivot
			right = append(right, element)
		} else {
			// - equal to pivot
			equalToPivot = append(equalToPivot, element)
		}
	}

	// return the concatenation of quickSort(less) + pivot + quickSort(larger)
	return append(quickSort(left), append(equalToPivot, quickSort(right)...)...)

}

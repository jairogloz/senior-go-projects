package util_test

import (
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/slices/remove_element/util"
	"testing"
)

func BenchmarkRemoveLoop(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	index := 2
	for i := 0; i < b.N; i++ {
		util.RemoveLoop(a, index)
	}
}

func BenchmarkRemoveNewSlice(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	index := 2
	for i := 0; i < b.N; i++ {
		util.RemoveNewSlice(a, index)
	}
}

func BenchmarkRemoveCopy(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	index := 2
	for i := 0; i < b.N; i++ {
		util.RemoveCopy(a, index)
	}
}

func BenchmarkRemoveAppend(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	index := 2
	for i := 0; i < b.N; i++ {
		util.RemoveAppend(a, index)
	}
}

func BenchmarkRemoveWithSlicesPackage(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	index := 2
	for i := 0; i < b.N; i++ {
		util.RemoveWithSlicesPackage(a, index)
	}
}

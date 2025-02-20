package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	s1 := "listen"
	s2 := "silent"

	s3 := "diavole in dracon limala asno"
	s4 := "leonardo davinci la monalisa"

	fmt.Println(s1, "|", s2, "=", isAnagram(s1, s2))
	fmt.Println(s3, "|", s4, "=", isAnagram(s3, s4))
	fmt.Println(s1, "|", s3, "=", isAnagram(s1, s3))

	fmt.Println(s1, "|", s2, "=", isAnagram2(s1, s2))
	fmt.Println(s3, "|", s4, "=", isAnagram2(s3, s4))
	fmt.Println(s1, "|", s3, "=", isAnagram2(s1, s3))
}

func isAnagram2(s1 string, s2 string) bool {
	s1 = strings.Replace(s1, " ", "", -1)
	s2 = strings.Replace(s2, " ", "", -1)
	if len(s1) != len(s2) {
		return false
	}

	s1Slice := []rune(s1)
	s2Slice := []rune(s2)
	slices.Sort(s1Slice)
	slices.Sort(s2Slice)

	for i, r := range s1Slice {
		if s2Slice[i] != r {
			return false
		}
	}

	return true
}

func isAnagram(s1 string, s2 string) bool {
	s1 = strings.Replace(s1, " ", "", -1)
	s2 = strings.Replace(s2, " ", "", -1)

	// Count element frequency
	s1Map := map[rune]int{}
	for _, i := range s1 {
		s1Map[i] = s1Map[i] + 1
	}
	s2Map := map[rune]int{}
	for _, i := range s2 {
		s2Map[i] = s2Map[i] + 1
	}
	for r, count := range s1Map {
		if s2Map[r] != count {
			return false
		}
	}

	return true
}

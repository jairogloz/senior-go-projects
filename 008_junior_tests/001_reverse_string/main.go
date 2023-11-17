package main

import "fmt"

func main() {
	// Reverse string
	// a,b,c,d
	// 0,1,2,3
	// i     j

	s := "abcd"

	fmt.Println(reverseString(s))
	fmt.Println(reverseStringPro(s))
}

func reverseStringPro(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func reverseString(s string) string {

	r := []rune(s)
	rReversed := make([]rune, len(s))
	// r = [a, b, c, d]
	// rReversed = [d,c, , ]

	j := len(r) - 1               //2
	for i := 0; i < len(r); i++ { //i=1
		rReversed[i] = r[j]
		j--
	}

	return string(rReversed)
}

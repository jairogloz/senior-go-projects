package main

import (
	"fmt"
	"reflect"
)

/*
	Letters and words
	Given two arrays one with letters and another with words, create  a function that returns the words that can be written with the letters available in the first array. Example

	letters = ["h", "e", "l", "o", "c", "f", "f", "x", "y", "e"]
	// h - 1
	// e - 2
	// l - 1
	words = ["hello", "coffee", "ley", "xyz"]
	// hello
	// h - 1 >
	// e- 1 >
	// l - 2 x
	Expected result: "coffee", "ley"
*/

func wordsFormable(letters []string, words []string) []string {
	// haz tu magia aquí!
	// contar letras en mi arreglo de letras
	letterCount := make(map[string]int)
	for _, letter := range letters {
		letterCount[letter] += 1
	}

	palabrasFormables := []string{}
forWords:
	for _, word := range words {
		// Contar las letras de cada palabra
		wordLetterCount := getWordLetterCount(word)
		//  map[e:1 h:1 l:2 o:1]
		for l, count := range wordLetterCount {
			if count > letterCount[l] {
				// no alcanzan
				// descartar palabra
				continue forWords
			}
		}

		// sí alcanzan las letras
		palabrasFormables = append(palabrasFormables, word)
	}
	return palabrasFormables
}

func getWordLetterCount(word string) map[string]int {
	letterCount := make(map[string]int)
	// "h-e-l-l-o" -> int32//rune
	for _, l := range word {
		// convert rune -> string
		letterCount[string(l)] += 1
	}
	return letterCount
}

func main() {
	// Esto no hay que modificarlo
	tests := []struct {
		letters  []string
		words    []string
		expected []string
	}{
		{[]string{"h", "e", "l", "o", "c", "f", "f", "x", "y", "e"}, []string{"hello", "coffee", "ley", "xyz"}, []string{"coffee", "ley"}},
		{[]string{"a", "b", "c"}, []string{"abc", "ab", "ac", "a"}, []string{"abc", "ab", "ac", "a"}},
		{[]string{"a", "b", "c", "d"}, []string{"aa", "aaa"}, []string{}},
	}

	for i, test := range tests {
		result := wordsFormable(test.letters, test.words)
		if reflect.DeepEqual(result, test.expected) {
			fmt.Printf("Test %d passed: %v\n", i+1, result)
		} else {
			fmt.Printf("Test %d failed: got %v, expected %v\n", i+1, result, test.expected)
		}
	}
}

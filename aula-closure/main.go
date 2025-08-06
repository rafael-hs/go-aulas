package main

import (
	"fmt"
)

func main() {

	checkR := CheckLetterOcurrence('r')

	count, word := checkR("Rafael Honorio da Silvarrr")

	fmt.Printf("Count of 'r': %d to word: %s \n", count, word)
}

func CheckLetterOcurrence(letter rune) func(word string) (int, string) {
	numberOfWords := 0
	return func(word string) (int, string) {
		count := 0
		wordSplited := []rune(word)

		for _, v := range wordSplited {
			if v == letter {
				count++
			}
		}
		numberOfWords++
		return count, word
	}
}

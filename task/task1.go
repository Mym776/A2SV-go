package main

import (
	"fmt"
	"strings"
	"unicode"
)

func palindrome(words string) bool {
	words = strings.ToLower(words)
	var word []rune
	for _, val := range words {
		if unicode.IsLetter(val) {
			word = append(word, val)
		}
	}

	start := 0
	end := len(word) - 1

	for start <= end {

		if word[start] != word[end] {
			return false
		}
		start += 1
		end -= 1
	}
	return true
}

func Counter(s string) map[string]int {

	if len(s) == 0 {
		return nil
	}

	s = strings.ToLower(s)
	var sentence string

	for _, val := range s {
		
		if unicode.IsLetter(val) || unicode.IsSpace(val) {
			sentence += string(val)
			
		}
	}
	
	var dictionary = make(map[string]int)
	words := strings.Split(sentence, " ")
	for _, val := range words {

		dictionary[val] += 1
	}
	return dictionary
}

func main() {
	
	// frequencey counter
	fmt.Println(Counter("GO is a staticaly typed language"))
	fmt.Println(Counter("I REALLY really like food"))
	fmt.Println(Counter("Are you sure? you do not look alright at all!"))
	
	// palindrome 
	fmt.Println(palindrome("hello"))
	fmt.Println(palindrome("atta"))
	fmt.Println(palindrome("AtTa"))
	fmt.Println(palindrome("At!@Ta"))

}

package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	word := ""
	len1, len2 := len(word1), len(word2)
	for i := 0; (i < len1) || (i < len2); i++ {
		if word1 != "" {
			word += word1[:1]
			word1 = word1[1:]
		}
		if word2 != "" {
			word += word2[:1]
			word2 = word2[1:]
		}
	}
	return word
}

func main() {
	word1 := "abc"
	word2 := "pqr"
	word := mergeAlternately(word1, word2)
	fmt.Println(word)
}

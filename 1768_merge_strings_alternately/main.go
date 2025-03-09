package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	word := ""
	for i := 0; (i < len(word1)) || (i < len(word2)); i++ {
		if word1 != "" {
			word1 = word1[0:]
			word += word1[:1]
		}
		if word2 != "" {
			word2 = word2[0:]
			word += word2[:1]
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
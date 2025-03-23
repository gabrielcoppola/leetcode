package main

import "fmt"

func find(pattern string, text string) int {
	window_size := len(pattern)
	text_len := len(text)

	for i := 0; i < (text_len - window_size + 1); i++ {
		if pattern == text[i:window_size+i] {
			return i
		}
	}
	return -1
}

func longestCommonPrefix(strs []string) string {
	base := strs[0]

	for i := 1; i < len(strs); i++ {
		for find(base, strs[i]) != 0 {
			if base != "" {
				base = base[:len(base)-1]
			} else {
				base = ""
			}
		}
	}
	return base
}

func main() {
	strs := []string{"c", "acc", "ccc"}
	result := longestCommonPrefix(strs)
	fmt.Println(result)

	// result := find("fl", "flower")
	// fmt.Println(result)

}

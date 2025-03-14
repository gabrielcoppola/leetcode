package main

import "fmt"

func isSubsequence(s string, t string) bool {
	a := ""
	idx_t := 0
	idx_s := 0
	for j := idx_s; j < len(s); j++ {
		for i := idx_t; i < len(t); i++ {
			if s[j] == t[i] {
				idx_s += 1
				idx_t = i
				a += string(s[j])
				break
			}
		}
		idx_t += 1
	}
	return a == s
}

func main() {
	s := "abc"
	t := "ahbgdc"
	boolean := isSubsequence(s, t)
	fmt.Println(boolean)
}
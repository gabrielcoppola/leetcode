package main

import "fmt"

func romanToInt(s string) int {
	teste := make([]int, 0)
	sum := 0
	dict := map[string]int {
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(s); i++ {
		for k, v := range dict {
			if string(s[i]) == k {
				teste = append(teste, v)
			}
		}
	}

	for i := 0; i < (len(teste)-1); i++ {
		if teste[i] < teste[i+1] {
			teste[i+1] = teste[i+1] - teste[i]
		} else {
			sum += teste[i]
		}
	}
	return (sum + teste[len(teste)-1])

}

func main() {
	str := "MCMXCIV"
	romanToInt(str)
}
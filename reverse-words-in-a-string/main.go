package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("abc def ghi"))
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	result := []string{}
	for _, v := range words {
		result = append(result, reverse(v))
	}
	return join(result)
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func join(words []string) string {
	joined := words[0]
	for i := 1; i < len(words); i++ {
		joined = joined + " " + words[i]
	}
	return joined
}

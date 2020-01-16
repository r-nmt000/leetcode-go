package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

func main() {
	fmt.Println(WordCount("test test test aaa aaa bbb"))

}

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	m := make(map[string]int)

	for _, word := range words {
		num, isExist := m[word]
		if isExist {
			m[word] = num + 1
		} else {
			m[word] = 1
		}
	}
	return m
}

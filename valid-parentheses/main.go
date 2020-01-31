package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValid("]"))
}

type stack []rune

func (s *stack) Push(v rune) {
	*s = append(*s, v)
}

func (s *stack) Pop(v rune) rune {
	if len(*s) == 0 {
		return '_'
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func isValid(s string) bool {
	pair := make(map[rune]rune)
	pair[')'] = '('
	pair[']'] = '['
	pair['}'] = '{'

	st := stack{}
	for _, c := range s {
		if validLeft(c) {
			st.Push(c)
		} else if validRight(c) {
			left := st.Pop(c)
			if left != pair[c] {
				return false
			}
		} else {
			// invalid rune
			return false
		}
	}

	if len(st) == 0 {
		return true
	} else {
		return false
	}
}

func validLeft(c rune) bool {
	validLeftRune := "({["
	return strings.ContainsRune(validLeftRune, c)
}

func validRight(c rune) bool {
	validRightRune := ")]}"
	return strings.ContainsRune(validRightRune, c)
}

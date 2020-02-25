package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("LLLRRRLRLRLLRR"))
}

func balancedStringSplit(s string) int {
	count := 0
	balanced := 0
	for _, v := range s {
		if v == 'R' {
			balanced++
		} else if v == 'L' {
			balanced--
		}

		if balanced == 0 {
			count++
		}
	}
	return count
}

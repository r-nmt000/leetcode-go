package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()

}

func numJewelsInStones(J string, S string) int {
	count := 0
	for _, c := range J {
		count += strings.Count(S, string(c))
	}
	return count
}

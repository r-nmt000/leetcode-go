package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}

func shortestToChar(S string, C byte) []int {
	N := len(S)
	ans := []int{}
	prev := math.MinInt32

	for i, v := range S {
		if byte(v) == C {
			prev = i
		}
		ans = append(ans, i-prev)
	}

	prev = math.MaxInt32
	for i := N-1; i >= 0; i-- {
		if byte(S[i]) == C {
			prev = i
		}
		ans[i] = min(ans[i], prev-i)
	}
	return ans

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
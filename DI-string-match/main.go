package main

import "fmt"

func main() {
	fmt.Println(diStringMatch("DIDI"))

}

func diStringMatch(S string) []int {
	N := len(S)
	lo, hi := 0, N
	ans := []int{}
	for i := 0; i < N; i++ {
		if string(S[i]) == "I" {
			ans = append(ans, lo)
			lo++
		} else {
			ans = append(ans, hi)
			hi--
		}
	}
	ans = append(ans, lo)
	return ans
}

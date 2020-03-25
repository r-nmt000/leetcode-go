package main

import "fmt"

func main() {
	fmt.Println(fib(10))

}

func fib(N int) int {
	if N == 0 { return 0 }
	if N == 1 { return 1 }
	memo := make(map[int]int)
	memo[0] = 0
	memo[1] = 1
	for i := 2;i <= N;i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[N]
}


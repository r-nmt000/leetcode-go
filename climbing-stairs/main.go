package main

import "fmt"

func main() {
	fmt.Println(climbStairsWithMemoRecursive(6))
	fmt.Println(climbStairsRecursive(6))
	fmt.Println(climbStairsDP(6))
}

func climbStairsRecursive(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return climbStairsRecursive(n-1) + climbStairsRecursive(n-2)
}

func climbStairsWithMemoRecursive(n int) int {
	memo := map[int]int{1: 1, 2: 2}
	climbStairsWithMemo(n, memo)
	return memo[n]
}

func climbStairsWithMemo(n int, memo map[int]int) int {
	if step, ok := memo[n]; ok {
		return step
	}
	memo[n] = climbStairsWithMemo(n-1, memo) + climbStairsWithMemo(n-2, memo)
	return memo[n]
}

func climbStairsDP(n int) int {
	memo := map[int]int{1: 1, 2: 2}
	for i := 3; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

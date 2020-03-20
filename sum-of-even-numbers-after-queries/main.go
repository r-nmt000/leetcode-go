package main

import "fmt"

func main() {
	A := []int{1,2,3,4}
	queries := [][]int{{1,0}, {-3,1}, {-4,0}, {2,3}}
	fmt.Println(sumEvenAfterQueries(A, queries))
}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	result := []int{}
	tmp := 0
	for i := 0; i < len(queries); i++ {
		val := queries[i][0]
		index := queries[i][1]
		A[index] += val
		for j := 0; j < len(A); j++ {
			if A[j] % 2 == 0 {
				tmp += A[j]
			}
		}
		result = append(result, tmp)
		tmp = 0
	}
	return result
}

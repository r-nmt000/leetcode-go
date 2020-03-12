package main

import "fmt"

func main() {
	row1 := []int{1, 1, 1, 1, 1, 1}
	row2 := []int{1, 1, 1, 1, 1, 1}
	row3 := []int{1, 1, 1, 1, 1, 1}
	//row1 := []int{1,1,0,0,0}
	//row2 := []int{1,1,1,1,0}
	//row3 := []int{1,1,0,0,0}
	//row4 := []int{1,1,1,1,1}
	mat := [][]int{row1, row2, row3}
	fmt.Println(kWeakestRows(mat, 1))

}

func kWeakestRows(mat [][]int, k int) []int {
	tmp := make([]int, len(mat))
	result := make([]int, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			tmp[i] += mat[i][j]
		}
	}
	for i := 0; i < len(mat[0])+1; i++ {
		for j, v := range tmp {
			if i == v {
				result = append(result, j)
			}
		}

	}
	return result[:k]
	//return tmp
}

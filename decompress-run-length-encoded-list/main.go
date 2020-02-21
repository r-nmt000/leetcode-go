package main

import "fmt"

func main() {
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))
}

func decompressRLElist(nums []int) []int {
	decompress := make([]int, 0)
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			decompress = append(decompress, nums[i+1])
		}
	}
	return decompress
}

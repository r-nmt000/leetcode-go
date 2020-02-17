package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 2, 2, 2, 3, 3, 4, 4}))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

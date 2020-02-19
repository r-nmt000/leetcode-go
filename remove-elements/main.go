package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{}, 3))
}

func removeElement(nums []int, val int) int {
	i := 0

	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

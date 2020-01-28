package main

import (
	"errors"
	"fmt"
)

func main() {
	ret, err := twoSumHashTable([]int{1, 2, 3, 4, 5, 6}, 7)
	fmt.Println(ret, err)
}

func twoSumBruteForce(nums []int, target int) ([]int, error) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == (nums[i] + nums[j]) {
				return []int{nums[i], nums[j]}, nil
			}
		}
	}
	return nil, errors.New("there is no answer")
}

func twoSumHashTable(nums []int, target int) ([]int, error) {
	table := make(map[int]int)
	for i, v := range nums {
		complement := target - v
		if _, ok := table[complement]; ok {
			return []int{table[complement], i}, nil
		}
		table[v] = i
	}
	return nil, errors.New("there is no answer")
}

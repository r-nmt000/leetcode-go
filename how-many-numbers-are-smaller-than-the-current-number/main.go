package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{1,2,4}
	//sortedNums := make([]int, len(nums))
	//copy(sortedNums, nums)
	//nums[0] = 3
	//fmt.Println(sortedNums)
	fmt.Println(smallerNumbersThanCurrent([]int{3, 5, 1, 3, 9}))
}

func smallerNumbersThanCurrent(nums []int) []int {
	result := []int{}
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Sort(sort.IntSlice(sortedNums))
	for _, v := range nums {
		result = append(result, firstIndexOf(v, sortedNums))
	}
	return result
}

func firstIndexOf(target int, ary []int) int {
	for i, v := range ary {
		if target == v {
			return i
		}
	}
	return -1
}

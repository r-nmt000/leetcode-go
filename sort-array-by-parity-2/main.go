package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity2([]int{4,2,5,7}))
}

func sortArrayByParity2(A []int) []int {
	for i := 0; i < len(A); i++ {
		if i%2 == 0 && A[i]%2 != 0 {
			A = swapElements(A, i, true)
		} else if i%2 == 1 && A[i]%2 != 1 {
			A = swapElements(A, i, false)
		}
	}
	return A
}

func swapElements(nums []int, startIndex int, isEven bool) []int {
	for i := startIndex; i < len(nums); i++ {
		if isEvenNum(nums[i]) == isEven {
			tmp := nums[startIndex]
			nums[startIndex] = nums[i]
			nums[i] = tmp
			break
		}
	}
	return nums
}

func isEvenNum(num int) bool {
	if num % 2 == 0 {
		return true
	}
	return false
}

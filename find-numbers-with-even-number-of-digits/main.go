package main

import "fmt"

func main() {
	fmt.Println(findNumbers([]int{12, 345, 6, 7896}))
	//fmt.Println(9/10)
}

func findNumbers(nums []int) int {
	count := 0
	for _, v := range nums {
		if calcDigits(v)%2 == 0 {
			count++
		}
	}
	return count
}

func calcDigits(num int) int {
	if num == 0 {
		return 1
	}
	digit := 0
	for num != 0 {
		digit++
		num /= 10
	}
	return digit
}

package main

import "fmt"

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))

}

func replaceElements(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}
	if len(array) == 1 {
		return []int{-1}
	}
	maxIndex, maxVal := findMax(array)
	for i := 0; i < maxIndex; i++ {
		array[i] = maxVal
	}
	return append(array[0:maxIndex], replaceElements(array[maxIndex:])...)
}

func findMax(array []int) (int, int) {
	maxIndex := 0
	maxVal := 0
	for i := 1; i < len(array); i++ {
		if array[i] >= maxVal {
			maxVal = array[i]
			maxIndex = i
		}
	}
	return maxIndex, maxVal
}

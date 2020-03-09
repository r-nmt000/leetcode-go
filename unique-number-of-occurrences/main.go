package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2}))

}

func uniqueOccurrences(arr []int) bool {
	hash := make(map[int]int)
	values := make([]int, 0)
	for i, v := range arr {
		if _, ok := hash[v]; ok {
			continue
		}
		count := countOccurrence(arr[i+1:], v)
		if contains(values, count) {
			return false
		}
		values = append(values, count)
		hash[v] = count
	}
	return true
}

func contains(values []int, target int) bool {
	for _, v := range values {
		if v == target {
			return true
		}
	}
	return false
}

func countOccurrence(arr []int, v int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if v == arr[i] {
			count++
		}
	}
	return count
}

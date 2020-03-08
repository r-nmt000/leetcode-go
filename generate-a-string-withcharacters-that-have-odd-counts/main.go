package main

import "fmt"

func main() {
	fmt.Println(generateTheString(12))
}

func generateTheString(n int) string {
	result := ""
	for i := 0; i < n-1; i++ {
		result = result + "a"
	}
	if (n % 2) == 1 {
		result = result + "a"
	} else {
		result = result + "b"
	}
	return result
}

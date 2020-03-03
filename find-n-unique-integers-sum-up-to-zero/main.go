package main

func main() {

}

func sumZero(n int) []int {
	isOdd := n%2 == 1
	result := []int{}
	for i := 1; i <= n/2; i++ {
		result = append(result, i)
		result = append(result, i*-1)
	}
	if isOdd {
		result = append(result, 0)
	}
	return result
}

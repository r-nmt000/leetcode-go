package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(123))
}

func numberOfSteps(num int) int {
	return countSteps(num, 0)
}

func countSteps(num int, steps int) int {
	if num == 0 {
		return steps
	}
	if (num % 2) == 0 {
		num /= 2
	} else {
		num -= 1
	}
	steps++
	return countSteps(num, steps)
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(substractProductAndSum(234))

}

func substractProductAndSum(n int) int {
	product := 1
	sum := 0

	for digit := n % 10; n != 0; digit = n % 10 {
		product *= digit
		sum += digit
		n = n / 10
	}
	return product - sum
}

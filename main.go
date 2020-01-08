package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 1; i <= 10; i++ {
		prev := z
		z -= (z*z - x) / (2 * z)
		fmt.Println("count: ", i)
		fmt.Println("z: ", z)
		if (prev - z) == 0 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println("result", Sqrt(100))
}

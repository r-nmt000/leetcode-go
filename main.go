package main

import "fmt"

var i, j int = 1, 2

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}

func add(x int, y int) int {
	return x + y
}

func swap(x, y, z string) (string, string, string) {
	return y, x, z
}

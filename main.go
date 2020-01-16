package main

import "fmt"

func fibonacci() func() int {
	step1 := 0
	step2 := 1
	return func() int {
		ret := step1 + step2
		step1 = step2
		step2 = ret
		return ret
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

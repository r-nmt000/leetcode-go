package main

func main() {

}


func smallestRangeI(A []int, K int) int {
	min := A[0]
	max := A[0]
	for _,v := range A {
		min = small(min, v)
		max = large(max, v)
	}
	return large(0, max - min - 2*K)

}

func small( a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func large(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(maximum69Number(9669))
}

func maximum69Number(num int) int {
	s := strconv.Itoa(num)
	changed := strings.Replace(s, "6", "9", 1)
	maxNum, _ := strconv.Atoi(changed)

	return maxNum
}

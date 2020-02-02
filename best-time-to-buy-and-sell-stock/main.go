package main

import (
    "fmt"
    "math"
)
func main() {
    fmt.Println(maxProfit([]int{}))

}

func maxProfit(prices []int) int {
    bottom := math.MaxInt32
    maxProfit := 0

    for i := 0; i < len(prices); i++ {
        if prices[i] < bottom {
            bottom = prices[i]
        } else if prices[i] - bottom > maxProfit {
            maxProfit = prices[i] - bottom
        }
    }
    return maxProfit
}

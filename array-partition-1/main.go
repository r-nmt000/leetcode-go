package main

import (
    "sort"
    "fmt"
)

func main() {
    fmt.Println(arrayPairSum([]int{1,2,3,4,5,6}))

}

func arrayPairSum(nums []int) int {
    sort.Sort(sort.IntSlice(nums))
    result := 0
    for i, v := range nums {
        if (i%2) == 0 {
            result += v
        }
    }
    return result
}

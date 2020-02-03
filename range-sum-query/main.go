package main

import "fmt"

type NumArray struct {
    nums []int
    sums []int
}

func Constructor(nums []int) NumArray {
    sums := make([]int, len(nums))
    sums[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        sums[i] = sums[i-1] + nums[i]
    }
    return NumArray{nums: nums, sums: sums}
}

func (this *NumArray) SumRangeCache(i int, j int) int {
    return this.sums[j] - this.sums[i-1]
}

func (this *NumArray) SumRangeOnTime(i int, j int) int {
    result := 0
    for ;i <= j; i++ {
        result += this.nums[i]
    }
    return result
}

func main() {
    obj := Constructor([]int{1,2,3,4,5,6,7})
    fmt.Println(obj.SumRangeCache(1,3))
}

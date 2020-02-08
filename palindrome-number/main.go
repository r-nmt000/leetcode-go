package main

import "fmt"

func main() {
    fmt.Println(isPalindrome(10))
}


func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    array := []int{}
    for ; x >= 10; x /= 10 {
        array = append(array, x % 10)
        fmt.Println(array)
        fmt.Println(x)
    }
    array = append(array, x % 10)
    fmt.Println(array)
    for i := 0; i < len(array); i++ {
        if array[i] != array[(len(array) - 1) - i] {
            return false
        }
    }
    return true
}

package main

import "fmt"

func main() {
    fmt.Println(judgeCircle("UUDDUUDDRL"));
}

func judgeCircle(moves string) bool {
  upCount := 0
  downCount := 0
  leftCount := 0
  rightCount := 0
  for _, v := range moves {
    switch v {
      case 'U':
        upCount++
      case 'D':
        downCount++
      case 'L':
        leftCount++
      case 'R':
        rightCount++
      default:
        return false
    }
  }
  if (upCount == downCount) && (leftCount == rightCount) {
      return true
  } else {
      return false
  }
}


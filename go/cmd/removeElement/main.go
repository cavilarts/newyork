package main

import (
  "fmt"
)

func removeElement(nums []int, val int) int {
  if len(nums) == 1 && nums[0] == val {
    return 0
  }

  if len(nums) == 1 && nums[0] != val {
    return 1
  }

  i := 0
  l := len(nums) - 1
 
  for i < l {
    if nums[l] == val {
      l--
    }

    if nums[i] == val {
      temp := nums[i]
      nums[i] = nums[l]
      nums[l] = temp
    }
    
    i++
  }

  fmt.Printf("%v | %v", i, l)

  return i
}

func main()  {
  ls := []int{3,3}
  re := removeElement(ls, 3)

  fmt.Printf("%+v", ls[:re])
}

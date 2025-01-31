package main

import (
  "fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
  k := m + n - 1
  i := m - 1
  j := n - 1

  for j >= 0 {
    if i >= 0 && nums1[i] > nums2[j] {
      nums1[k] = nums1[i]
      i--
    } else {
      nums1[k] = nums2[j]
      j--
    }
    k--
  }
}

func main()  {
  n1 := []int{1,2,3,0,0,0}
  n2 := []int{2,5,6}

  merge(n1, 3, n2, 3)
  
  fmt.Printf("result: %+v", n1)
}

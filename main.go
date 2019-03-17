package main

import (
	"fmt"
	"algorithm/leetCode"
)
func main()  {
	nums := []int { 2, 7, 11, 15}
	target := 9
	arr := leetCode.TwoSum(nums,target)
	fmt.Println(arr)
}
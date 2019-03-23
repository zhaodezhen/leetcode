package main

import (
	"algorithm/leetCode"
	"fmt"
)

func main() {
	/* 两数之和
	nums := []int { 2, 7, 11, 15}
	target := 9
	arr := leetCode.TwoSum(nums,target)
	fmt.Println(arr)
	*/
	/* 只出现一次的数 */
	nums := []int{2, 2, 1}
	res := leetCode.SingleNumber(nums)
	fmt.Println(res)
}

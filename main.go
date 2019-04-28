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

	/* 只出现一次的数
	nums := []int{2, 2, 1}
	result := leetCode.SingleNumber(nums)
	fmt.Println(result)
	*/

	/* 整数反转
	x := 1002
	result := leetCode.Reverse(x)
	fmt.Println(result)
	 */

	/* 爬楼梯
	result := leetCode.ClimbStairs(6)
	fmt.Println(result)
	 */

	/* Maximum Subarray
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	result := leetCode.MaxSubArray(nums)
	fmt.Println(result)
	*/

	/*	Best Time to Buy and Sell Stock 1 & 2
	prices := []int{6,1,3,2,4,7}
	result := leetCode.MaxProfit2(prices)
	fmt.Println(result)
	*/
	/* Min Cost Climbing Stairs  */
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	result := leetCode.MinCostClimbingStairs(cost)
	fmt.Println(result)
}

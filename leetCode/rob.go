package leetCode

/*
	You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

	Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

	Example 1:

	Input: [1,2,3,1]
	Output: 4
	Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
				 Total amount you can rob = 1 + 3 = 4.
	Example 2:

	Input: [2,7,9,3,1]
	Output: 12
	Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
	 			 Total amount you can rob = 2 + 9 + 1 = 12.
	典型的动态规划，看了两遍题就开始做了，期初的思路是 nums[i-1]+nums[i+1] 这种思路，求每隔一个数的大小
	主要代码:
	if n % 2 == 0 {
		tow += nums[i]
	} else {
		one += nums[i]
	}
	但是运行了没有成功，因为会有 [2,1,1,2] 这种情况下上面的代码结果是 3 ，但是要求 是 2 + 2，也就是找出最大的，
但是之间的间隔不限于 1 。
	换一个思路，记录两个变量 previous、current 分别记录一下之前的值，和当前最大的，创建一个临时变量保存 current
判断 previous + nums[i] 和 current 谁大，并赋值给 current， previous 等于赋值前的 current （临时变量）,一句话概括；在满足不触动警报装置（就是不取连续数据）的情况下，取的{一个}最大的数进行计算。

	时间复杂度: O(n)
	空间复杂度: O(1)
	Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber.
	Memory Usage: 2 MB, less than 100.00% of Go online submissions for House Robber.
*/
func Rob(nums []int) int {
	l := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	if l < 1 {
		return 0
	}else if l < 2 {
		return nums[0]
	}

	previous := nums[0]
	current := max(nums[0], nums[1])
	for i := 2; i < l; i++ {
		temp := current
		current = max(previous+nums[i], current)
		previous = temp;
	}
	return current
}

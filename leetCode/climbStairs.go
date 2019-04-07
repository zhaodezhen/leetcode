package leetCode

/*
	You are climbing a stair case. It takes n steps to reach to the top.

	Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

	Note: Given n will be a positive integer.

	Example 1:

	Input: 2
	Output: 2
	Explanation: There are two ways to climb to the top.
	1. 1 step + 1 step
	2. 2 steps
	Example 2:

	Input: 3
	Output: 3
	Explanation: There are three ways to climb to the top.
	1. 1 step + 1 step + 1 step
	2. 1 step + 2 steps
	3. 2 steps + 1 step
*/
/*
	爬楼梯问题，经典的动态规划题
	首先这个题型有点像斐波那契数列，可以考虑先用递归实现一遍
func ClimbStairs(n int) int {
	if n < 2 {
		return 1
	}
	return ClimbStairs(n-1) + ClimbStairs(n-2)
}
	提交 Time Limit Exceeded ，看来递归不行，当N很大时复杂度就会急速增加。时间复杂度为O(2^n)
	从上面的解法中可以看到，要解这个问题，需要将其拆成若干部分（就是子问题），再根据子问题的解得出原问题的解。
	这恰恰是动态规划的核心思想。所以用动态规划试试(这个理解起来可能有点难，可以把它倒序为斐波那契数列来计算)。
	一个人到达N层楼梯包括两种方法（N为楼层J为步数）：
	1.当J=1时，N-J，在爬一步到
	2.当J=2时，N-J，在爬两步到
	所以根据这个得到总方法为 I = N-J + N-J
func ClimbStairs(n int) int {
	if n < 2 {
		return n
	}
	res := make([]int, n+1)
	res[0] = 1
	res[1] = 1
	for i := 2; i <= n; i++ {
		one := res[i-1]
		tow := res[i-2]
		res[i] = one + tow
	}
	return res[n]
}
	上面这个解法，时间复杂度为O(n)，空间复杂度也为O(n)
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
	Memory Usage: 2 MB, less than 100.00% of Go online submissions for Climbing Stairs.
	看起来不错，但是结果没必要存一个数组，进一步优化内存，我们只依赖两种方法，所以可以考虑用两个变量存储方法一个存储结果：
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
	Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Climbing Stairs.
*/
func ClimbStairs(n int) int {
	if n < 4 {
		return n
	}
	var a uint32 = 2
	var b uint32 = 3
	var ret uint32 = 5

	for i := 5; i <= n; i++ {
		a = ret
		ret = b + ret
		b = a
	}
	return int(ret);
}

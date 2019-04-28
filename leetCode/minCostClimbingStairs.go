package leetCode

/*
	On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).

	Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.

	Example 1:
	Input: cost = [10, 15, 20]
	Output: 15
	Explanation: Cheapest is start on cost[1], pay that cost and go to the top.
	Example 2:
	Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
	Output: 6
	Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].
	Note:
	cost will have a length in the range [2, 1000].
	Every cost[i] will be an integer in the range [0, 999].
*/
/*
	和爬楼梯的题有点类似，一次可以走一步或者两步，求最短路径
	Runtime: 4 ms, faster than 100.00% of Go online submissions for Min Cost Climbing Stairs.
	Memory Usage: 3.1 MB, less than 30.77% of Go online submissions for Min Cost Climbing Stairs.
	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func MinCostClimbingStairs(cost []int) int {
	l := len(cost)
	dp := make([]int, l)
	min := func(a, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}
	dp[0],dp[1] = cost[0],cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
	}
	return min(dp[l-1], dp[l-2])
}


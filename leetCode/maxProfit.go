package leetCode

/*
	Say you have an array for which the ith element is the price of a given stock on day i.

	If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

	Note that you cannot sell a stock before you buy one.

	Example 1:

	Input: [7,1,5,3,6,4]
	Output: 5
	Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
				 Not 7-1 = 6, as selling price needs to be larger than buying price.
	Example 2:

	Input: [7,6,4,3,1]
	Output: 0
	Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/
/*
	买卖股票最佳时机，动态规划类型，
	感觉做自己动态规划的题的时候推导公式这方面很弱，连着做了几道这种套路的题了，但是做这个题的时候还是用了半个多小时。
	思路倒是可以跟上了，但是验证思路的推导公式的时候还是很弱，提交了四次才通过。
	注意：多读几遍题，不要上来就做，不然很容易出错！先找到最便宜的时机买，然后最贵的时候抛出。
	卖出一定在买入后，不能先卖在卖。
	注意边界问题，防止索引超出范围！
	思路，记录三个数，最佳买入，和最佳卖出，和收益。
	判断当前价格是不是最低的，如果是就记录为最佳买入时间，并初始化最佳卖出时间，因为卖出是在买入之后（最低买入的不一定带来最高的收益）
	判断当前价格是不是最高的，如果是就记录为最佳卖出时间
	判断当前的收益是不是最大的，如果是则记录

	步骤：
	1. 初始化买入、卖出价格，和最佳收益
	循环;
	2. 比较当前价格是不是最低的，如果是就更新买入价格，因为卖出在买入后，所以也要更新卖出价格。
	3. 比较当前价格是不是最高，如果是就更新卖出价格
	4. 比较当前买入和卖出是不是最佳的收益，如果是则记录

func MaxProfit(prices []int) int {
	l := len(prices)
	if l < 1 {
		return 0
	}
	min := prices[0]
	max := min
	ret := 0
	for i := 1; i < l; i++ {
		if min > prices[i] {
			min = prices[i]
			max = min
		}
		if max < prices[i] {
			max = prices[i]
		}
		if ret < max-min {
			ret = max - min
		}
	}
	return ret
}
	时间复杂度: O(n)
	空间复杂度: O(1)
	Runtime: 4 ms, faster than 100.00% of Go online submissions for Best Time to Buy and Sell Stock.
	Memory Usage: 3.1 MB, less than 35.56% of Go online submissions for Best Time to Buy and Sell Stock.

	代码效率没问题，下面看看有没有内存空间需要优化的地方。
	我们仔细看代码，其实最佳完全可以不用，只需要 最佳买入和收益就行了。下面是优化过的代码；
	Runtime: 4 ms, faster than 100.00% of Go online submissions for Best Time to Buy and Sell Stock.
	Memory Usage: 3.1 MB, less than 68.89% of Go online submissions for Best Time to Buy and Sell Stock.

func MaxProfit(prices []int) int {
	l := len(prices)
	if prices == nil || l < 1 {
		return 0
	}
	min := prices[0]
	ret := 0
	for i := 1; i < l; i++ {
		if min > prices[i] {
			min = prices[i]
		}
		if ret < prices[i]-min {
			ret = prices[i] - min
		}
	}
	return ret
}

*/
/*
	Say you have an array for which the ith element is the price of a given stock on day i.

	Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

	Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).
*/
/*
	第一道题的变形，唯一区别是可以多次买卖，就是说下个数如果是跌，我就把当前的数卖出，然后下个低的数在买入
	判断我本次买下次是否能撑赚，不赚就不买，赚就买。
	做的时候一直纠结动态规划，其实这是一个贪心算法，
	总是做出在当前看来是最好的选择，不从整体最优上加以考虑，也就是说，只关心当前最优解
	时间复杂度: O(n)
	空间复杂度: O(1)
	Runtime: 4 ms, faster than 100.00% of Go online submissions for Best Time to Buy and Sell Stock II.
	Memory Usage: 3.1 MB, less than 75.00% of Go online submissions for Best Time to Buy and Sell Stock II.
*/
func MaxProfit2(prices []int) int {
	l := len(prices)
	ret := 0
	if prices == nil || l < 1 {
		return 0
	}
	for i := 0; i < l-1; i++ {
		if prices[i] < prices[i+1] {
			ret += prices[i+1] - prices[i]
		}
	}
	return ret
}

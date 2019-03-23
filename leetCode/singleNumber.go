/*
	Given a non-empty array of integers, every element appears twice except for one. Find that single one.

	Note:

	Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

	Example 1:

	Input: [2,2,1]
	Output: 1
	Example 2:

	Input: [4,1,2,1,2]
	Output: 4
*/

/*
   这个题比较简单，其中有一个出现一次，需要把它找出来。其余的全部都出现两次。
   要求是不适用额外空间，时间复杂度是线性
   这个题首先想到的是给数组排序，然后进行比较当前的数和下一个数，但是这种解法的查找和空间并不符合答案的要求
   最终采用了异或运算，这个比较方便，也不会占用额外的空间
   时间复杂度: O(n) 遍历了一次数组
   空间复杂度: O(1) 记录了一个 res
   Runtime: 8 ms
   Memory Usage: 4.7 MB
*/

package leetCode

func SingleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

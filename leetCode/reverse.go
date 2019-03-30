package leetCode

import "math"

/*
	Given a 32-bit signed integer, reverse digits of an integer.

	Example 1:
	```
	Input: 123
	Output: 321
	```
	Example 2:
	```
	Input: -123
	Output: -321
	```
	Assume we are dealing with an environment which could only store integers within the 32-bit
	signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your
	function returns 0 when the reversed integer overflows.
*/

/*
	整数反转
	采用取模运算每次获取第一位的数，然赋值给num。
	其中需要注意的有两点，一是溢出问题，二是为0的数。
	所以采用num*10加上当前取模得数，并判断整数是否溢出
    时间复杂度: O(n)
	空间复杂度: O(1)
	Runtime: 4 ms
	Memory Usage: 2.2 MB
*/
func Reverse(x int) int {
	num := 0
	for x != 0 {
		num = num*10 + x%10
		x = x / 10
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	return num

}

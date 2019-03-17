package leetCode

/*
	Given an array of integers, return indices of the two numbers such that they add up to a specific target.

	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	Example:

	Given nums = [2, 7, 11, 15], target = 9,

	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].
*/

/*
	方法1暴力法
	两层循环逐一比对，得到两数之和为目标值则成功。
	时间复杂度: O(n^2)
	空间复杂度: O(1)
	Runtime: 52 ms
	Memory : 2.9 MB

	func TwoSum(nums []int, target int) []int {
		l := len(nums)
		result := make([]int, 2)
		for i := 0; i < l; i++ {
			for j := l - 1; j > i; j-- {
				if nums[i]+nums[j] == target {
					result[0] = i
					result[1] = j
				}
			}
		}
		return result
	}

*/
/*
	方法2hash法
	首先我们记录一个hash，以空间换时间，我们可以将查找时间从O(n)降低到 O(1)
	迭代中 我们取 target−num  判断是否存在hash表中，如果存在则查找成功，直接返回表中下标和当前下标。
	如果不存在，则保存本次的值，继续执行
	时间复杂度: O(n) 只遍历了一次数组，查找花费O(1)
	空间复杂度: O(n) 最坏情况下需要记录n个数组
	Runtime: 4 ms
	Memory : 3.7 MB
*/

func TwoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		diff := target - num
		//获取其中一个数的值
		if j, ok := m[diff]; ok { //判断哪个数是否存在
			return []int{j, i}
		}
		m[num] = i //不存在则把当前数记录下，供下次循环查找

	}
	return []int{}

}


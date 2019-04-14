package leetCode

/*
	Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

	Example:

	Input: [-2,1,-3,4,-1,2,1,-5,4],
	Output: 6
	Explanation: [4,-1,2,1] has the largest sum = 6.
	Follow up:

	If you have figured out the O(n) solution, try coding another solution using the divide
	and conquer approach, which is more subtle.
*/

/*
	最大子序和，可以用动态规划做也可以用分治法做，这里采用动态规划
	max 为连续最大和 ，sum为当前记录的连续数据，如果是 sum 为负数则从新为sum赋值
	e.g.
	-2,1,-3,4,-1,2,1,-5,4
    -2						sum = -2 ; max = -2
    -2,1					sum = 1  ; max = 1
	-2,1,-3					sum = -2 ; max = 1
    -2,1,-3,4 				sum = 4  ; max = 4
	-2,1,-3,4,-1			sum = 3  ; max = 4
	-2,1,-3,4,-1,2			sum = 5  ; max = 4
	-2,1,-3,4,-1,2,1		sum = 6  ; max = 6
	-2,1,-3,4,-1,2,1,-5		sum = 1  ; max = 6
	-2,1,-3,4,-1,2,1,-5,4	sum = 5  ; max = 6
	时间复杂度: O(n)
	空间复杂度: O(1)
	Runtime: 4 ms, faster than 100.00% of Go online submissions for Maximum Subarray.
	Memory Usage: 3.3 MB, less than 80.33% of Go online submissions for Maximum Subarray.
*/

func MaxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	sum := nums[0]
	max := sum

	for i := 1; i < l; i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum = sum + nums[i]
		}
		if max < sum {
			max = sum
		}

	}
	return max
}

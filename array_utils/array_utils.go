package array_utils

import "fmt"

/**
求最大子序列和 （就是说子序列加起来和最大）
*/
func FindMaxSeqSum(array []int) int {
	seqSum := make([]int, 0) // 存储子序列和
	// 初始子序列和为 数组下标为0的值
	seqSum = append(seqSum, array[0])
	for i := 1; i < len(array); i++ {
		if array[i] > seqSum[i-1]+array[i] {
			seqSum = append(seqSum, array[i])
		} else {
			seqSum = append(seqSum, seqSum[i-1]+array[i])
		}
	}
	max := seqSum[0]
	for j := 1; j < len(seqSum); j++ {
		if seqSum[j] > seqSum[j-1] {
			max = seqSum[j]
		}
	}
	return max
}

/**
二分查找法
查找某个值在有序数组中是否存在
*/
func BinaryFindOrderArray(array []int, value int) bool {
	head := 0
	tail := len(array) - 1
	for head <= tail {
		mid := (head + tail) >> 1
		if array[mid] == value {
			return true
		} else if array[mid] > value {
			tail = mid - 1
		} else {
			head = mid + 1
		}
	}
	return false
}

/**
数组是有序的
在数组中查找匹配value的第一个下标位置
*/
func BinaryFindFirstOrderArray(array []int, value int) int {
	head := 0
	height := len(array) - 1
	for head <= height {
		mid := head + (height-head)>>1
		if value > array[mid] {
			head = mid + 1
		} else if value < array[mid] {
			height = mid - 1
		} else {
			if mid == 0 || array[mid-1] != value {
				return mid
			}
			height = mid - 1
		}
	}
	return -1
}

/**
查找有序数组中匹配目标的最后一个位置的下标
*/
func BinaryFindTailOrderArray(array []int, value int) int {
	head := 0
	tail := len(array) - 1
	for head <= tail {
		mid := head + (tail-head)>>1
		if array[mid] > value {
			tail = mid - 1
		} else if array[mid] < value {
			head = mid + 1
		} else {
			if mid == len(array)-1 || array[mid+1] != value {
				return mid
			}
			head = mid + 1
		}
	}
	return -1
}

/**
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
说明:
初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:
输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3
输出: [1,2,2,3,5,6]
*/
func MergeTwoArray(nums1 []int, m int, nums2 []int, n int) {
	if n > 0 {
		for i := 0; i < n; i++ {
			nums1[m+i] = nums2[i]
		}
	}
	lindex := 0
	rindex := m
	for lindex < m && len(nums1) > rindex {
		for lindex < m && nums1[lindex] > nums1[rindex] {
			nums1[lindex], nums1[rindex] = nums1[rindex], nums1[lindex]
			//使右边重新变得有序
			for (rindex + 1) < (m + n) {
				if nums1[rindex] < nums1[rindex+1] {
					break
				}
				nums1[rindex], nums1[rindex+1] = nums1[rindex+1], nums1[rindex]
				rindex++
			}
			rindex = m
		}
		lindex++
	}
	fmt.Println(nums1)

}

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:
如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/
/**
 * 定义状态：
 * dp[i] ： 表示以 nums[i] 结尾的连续子数组的最大和
 * <p>
 * 状态转移方程：
 * dp[i] = max{num[i],dp[i-1] + num[i]}
 *
 * @param nums
 * @return
 */
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for index, v := range nums {
		if index == 0 {
			dp[index] = v
		} else {
			if dp[index-1]+v > v {
				dp[index] = dp[index-1] + v
			} else {
				dp[index] = v
			}
		}
	}
	max := dp[0]
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}
/**
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
示例 2:

输入: [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例 3:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 */
func maxProfit(prices []int) int {

	return 0
}

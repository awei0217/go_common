package array_utils

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
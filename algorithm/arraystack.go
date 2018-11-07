package algorithm

import "strconv"

/**
数组实现栈
*/

type ArrayStack struct {
	Size int
	Data []interface{}
}

func Push(stack *ArrayStack, value interface{}) *ArrayStack {
	if stack == nil {
		stack = &ArrayStack{
			Size: 0,
			Data: make([]interface{}, 10, 10),
		}
		stack.Data[stack.Size] = value
	} else {
		if stack.Size == len(stack.Data) {
			stack.Data = append(stack.Data, value)
		} else {
			stack.Data[stack.Size] = value
		}
	}
	stack.Size++
	return stack
}

func Pop(stack *ArrayStack) interface{} {
	if stack == nil || stack.Size == 0 {
		return nil
	}
	value := stack.Data[stack.Size-1]
	stack.Data[stack.Size-1] = nil
	stack.Size--
	return value
}

func Peek(stack *ArrayStack) interface{} {
	if stack == nil {
		return nil
	}
	value := stack.Data[stack.Size-1]
	return value
}

func IsEmpty(stack *ArrayStack) bool {
	if stack == nil || stack.Data == nil {
		return true
	}
	if stack.Size == 0 {
		return true
	}
	return false
}

/**
给定两个没有重复元素的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。
nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出-1。
示例 1:
输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
输出: [-1,3,-1]
解释:
    对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
    对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
    对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。

示例 2:
输入: nums1 = [2,4], nums2 = [1,2,3,4].
输出: [3,-1]
解释:
    对于num1中的数字2，第二个数组中的下一个较大数字是3。
    对于num1中的数字4，第二个数组中没有下一个更大的数字，

*/
func NextGreaterElement(findNums []int, nums []int) []int {
	var stack *ArrayStack
	m := make(map[int]int)
	for _, v := range nums {
		if !IsEmpty(stack) && Peek(stack).(int) < v {
			for !IsEmpty(stack) {
				flag := true
				if Peek(stack).(int) < v {
					m[Pop(stack).(int)] = v
					flag = false
				}
				if flag {
					break
				}
			}
		}
		stack = Push(stack, v)
	}
	for index, v := range findNums {
		if _, ok := m[v]; ok {
			findNums[index] = m[v]
		} else {
			findNums[index] = -1
		}
	}
	return findNums
}

func CalPoints(ops []string) int {
	if len(ops) == 0 {
		return 0
	}
	var stack *ArrayStack
	for _, v := range ops {
		i := 0
		if v == "+" {
			if stack == nil || stack.Size == 0 {
				i = 0
			} else if stack.Size == 1 {
				i = Peek(stack).(int)
			} else {
				m := Pop(stack).(int)
				n := Peek(stack).(int)
				i = m + n
				stack = Push(stack, m)
			}
		} else if v == "D" {
			if stack == nil || stack.Size == 0 {
				i = 0
			} else {
				m := Peek(stack).(int)
				i = m * 2
			}
		} else if v == "C" {
			Pop(stack)
		} else {
			i, _ = strconv.Atoi(v)
		}
		if i != 0 {
			stack = Push(stack, i)
		}
	}
	sum := 0
	for !IsEmpty(stack) {
		m := Pop(stack).(int)
		sum += m
	}
	return sum
}

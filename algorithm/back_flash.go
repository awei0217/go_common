package algorithm

import (
	"fmt"
)

/**
回溯算法
*/
var result []int = make([]int, 8, 8)

//八皇后,result 下标表示行，值表示列
func Call8Queens(row int) {
	if row == 8 { // 8个棋子都放置好了
		printResult(result) //打印结果
		return
	}
	for column := 0; column < 8; column++ {
		if isOk(row, column, result) {
			result[row] = column
			Call8Queens(row + 1)
		}
	}
}

func isOk(row int, column int, result []int) bool {
	left, right := column-1, column+1
	for i := row - 1; i >= 0; i-- {
		if result[i] == column { // 第i行和这一行是同一列
			return false
		}
		if left >= 0 { //考察左对角线
			if result[i] == left {
				return false
			}
		}
		if right < 8 { //考察右对角线
			if result[i] == right {
				return false
			}
		}
		left--
		right++
	}
	return true
}

func printResult(result []int) {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if result[row] == column {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println()

}

/**
求数字的全排列 比如 1,2,3,4,5,6
array 一组数字  temp 其实下标
*/
func F(array []int, temp int) {
	if temp == len(array) {
		fmt.Println(array)
		return
	}
	for i := temp; i < len(array); i++ {
		array[temp], array[i] = array[i], array[temp]
		F(array, temp+1)
		array[temp], array[i] = array[i], array[temp]
	}
}

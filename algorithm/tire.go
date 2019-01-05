package algorithm

import "fmt"

type TireNode struct {
	num   int         // 有多少个单词通过这个节点
	son   []*TireNode // 子节点
	isEnd bool        // 是否最后一个节点
	val   rune        // 节点的值
}

const SIZE = 26

type TireTree struct {
	root *TireNode // 根节点
}

/**
字典树添加
*/
func (tt *TireTree) Insert(str string) {
	if len(str) == 0 {
		return
	}
	if tt.root == nil {
		tt.root = &TireNode{}
	}
	tn := tt.root
	for _, v := range str {
		index := v - 'a' // 计算字符的下标
		if tn.son == nil {
			tn.son = make([]*TireNode, SIZE)
		}
		if tn.son[index] == nil {
			tn.son[index] = &TireNode{
				val: v,
				num: 1,
			}
		} else {
			tn.son[index].num++
		}
		tn = tn.son[index]
	}
	tn.isEnd = true
}

/**
查找字符串是否存在
*/
func (tt *TireTree) Query(str string) bool {
	if tt.root == nil {
		return false
	}
	tn := tt.root
	for _, v := range str {
		index := v - 'a'
		if tn.son == nil {
			break
		}
		if tn.son[index] == nil || tn.son[index].val != v {
			break
		}
		tn = tn.son[index]
	}
	return tn.isEnd
}

// 前序遍历字典树.
func (tt *TireTree) PreTraverse(tn *TireNode) {
	if tn == nil {
		return
	}
	fmt.Print(string(tn.val))
	for i := 0; i < len(tn.son); i++ {
		tt.PreTraverse(tn.son[i])
	}
}

/**
统计以prefix开头的单词有多少个
*/
func (tt *TireTree) CountPrefix(prefix string) int {
	tn := tt.root
	if tn == nil {
		return 0
	}
	for _, v := range prefix {
		index := v - 'a'
		if tn.son[index] == nil {
			return 0
		}
		if tn.son[index].val != v {
			return 0
		}
		tn = tn.son[index]
	}
	return tn.num
}

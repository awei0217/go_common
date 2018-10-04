package algorithm

import (
	"errors"
	"fmt"
)

const(
	VALUE_REPEAT = "添加的值在二叉树中已经存在"
)

type AvlNode struct {
	value int // 节点的值
	left *AvlNode
	right *AvlNode
	parent *AvlNode
}
type AvlTree struct {
	// 树的跟节点
	root *AvlNode
}

/**
	创建一颗空树
 */
func CreateAvlTree()*AvlTree  {

	return &AvlTree{}
}
/**
	添加一个节点
 */
func (avlTree *AvlTree) Insert(value int)error {
	// 如果根节点为空，给节点赋值
	newNode := &AvlNode{value:value}
	if avlTree.root == nil {
		avlTree.root =newNode
		return nil
	}
	root := avlTree.root
	//添加节点
	for{
		if value < root.value{
			if root.left == nil{
				root.left =newNode
				newNode.parent = root
				break
			}
			root = root.left
		}else if value > root.value{
			if root.right == nil{
				root.right =newNode
				newNode.parent = root
				break
			}
			root = root.right
		}else{
			return errors.New(VALUE_REPEAT)
		}
	}
	//旋转
	for root != nil{
		if GetHeight(root.left) - GetHeight(root.right) == 2{
			if value < root.left.value{
				//左左旋转
				leftLeftRotate(root,avlTree)
			}else{
				//左右旋转
				leftRightRotate(root,avlTree)
			}
		}
		if GetHeight(root.right) - GetHeight(root.left) == 2{
			if value > root.right.value{
				//右右旋转
				rightRightRotate(root,avlTree)
			}else{
				//右左旋转
				rightLeftRotate(root,avlTree)
			}
		}
		root = root.parent
	}
	return nil
}


func GetHeight(childNode *AvlNode)int {
	if childNode == nil{
		return 0
	}
	temp  := childNode
	left := 1
	right := 1
	for temp.left != nil{
		left++
		temp = temp.left
	}
	temp = childNode
	for temp.right != nil{
		right++
		temp = temp.right
	}
	if left > right{
		return left
	}
	return right
}
/**
	左左旋转
 */
func leftLeftRotate(root *AvlNode,tree *AvlTree){
	p := root.parent
	l := root.left
	r := l.right
	l.right= root
	l.parent = p
	root.parent = l
	root.left = r
	if p != nil{
		p.left =l
	}
	if tree.root == root{
		tree.root = l
	}
}
/**
	左右旋转
 */
func leftRightRotate(root *AvlNode,tree *AvlTree){
	p := root.parent
	l := root.left
	lr := l.right
	lr.left = l; l.parent = lr;l.right = nil
	lr.right = root;root.parent =lr;root.left=nil
	lr.parent = p
	if p != nil{
		p.left = lr
	}
	if tree.root == root{
		tree.root = lr
	}
}
/**
	右右旋转
 */
func rightRightRotate(root *AvlNode, tree *AvlTree) {
	p:= root.parent
	r := root.right
	r.parent = p
	l:= r.left
	r.left = root
	root.parent = r
	root.right = l
	if p != nil{
		p.right = r
	}
	if tree.root == root{
		tree.root = r
	}
}
/**
	右左旋转
 */
func rightLeftRotate(root *AvlNode,tree *AvlTree){
	p := root.parent
	r := root.right
	rl := r.left
	rl.right = r; r.parent = rl;r.left = nil
	rl.left = root;root.parent =rl;root.right=nil
	rl.parent = p
	if p != nil{
		p.right = rl
	}
	if tree.root == root{
		tree.root = rl
	}
}
func PrePrint(root *AvlNode){
	if root == nil{
		return
	}
	fmt.Print(root.value,",")
	PrePrint(root.left)
	PrePrint(root.right)
}


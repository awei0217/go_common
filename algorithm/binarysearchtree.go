package algorithm

type BinarySearchTree struct {
	value      int
	leftChild  *BinarySearchTree
	rightChild *BinarySearchTree
}

//递归添加
func (tree *BinarySearchTree) RecursionAdd(value int) {
	if tree.value == 0 {
		tree.value = value
		tree.leftChild = &BinarySearchTree{}
		tree.rightChild = &BinarySearchTree{}
	} else {
		if value < tree.value {
			tree.leftChild.RecursionAdd(value)
		} else {
			tree.rightChild.RecursionAdd(value)
		}
	}
}

/**
非递归添加
*/
func (tree *BinarySearchTree) NotRecursionAdd(value int) {
	if tree.value == 0 {
		tree.value = value
		tree.leftChild = &BinarySearchTree{}
		tree.rightChild = &BinarySearchTree{}
	} else {
		current := tree
		for {
			if value < current.value {
				current = current.leftChild
			} else if value > current.value {
				current = current.rightChild
			} else {
				return
			}
			if current == nil {
				current = &BinarySearchTree{value, nil, nil}
				break
			}
		}
	}
}

func (tree *BinarySearchTree) Search(value int) *BinarySearchTree {
	current := tree
	if current.value == value {
		return current
	}
	for {
		if current == nil {
			return nil
		}
		if value < current.value {
			current = current.leftChild
		} else if value > current.value {
			current = current.rightChild
		} else {
			return current
		}
	}

}

func (tree *BinarySearchTree) Remove(value int) {
	t := tree.Search(value)
	if t == nil {
		panic("value is not exist")
	}
	if t.leftChild.value == 0 && t.rightChild.value == 0 {
		destroy(t)
	} else if t.leftChild.value != 0 && t.rightChild.value != 0 { // Both LeftChild and RightChild are not empty.
		// Get the min-element in its RightChild.
		min := t.rightChild.GetMin()
		// Replace its value with the min-element's value.
		t.value = min.value
		// Remove the min-element.
		destroy(min)
	} else { // LeftChild or RightChild is empty.
		// Replace it with its child which is not empty.
		if t.leftChild.value > 0 {
			t.replaceWith(t.leftChild)
		} else {
			t.replaceWith(t.rightChild)
		}
	}
}
func (tree *BinarySearchTree) replaceWith(t *BinarySearchTree) {
	tree.value = t.value
	tree.leftChild = t.leftChild
	tree.rightChild = t.rightChild
	destroy(t)
}
func (tree *BinarySearchTree) GetMin() *BinarySearchTree {
	if tree.leftChild.value == 0 {
		return tree
	} else {
		return tree.leftChild.GetMin()
	}
}
func destroy(tree *BinarySearchTree) {
	tree.value = 0
	tree.leftChild = nil
	tree.rightChild = nil
}

/**
获取树的深度
*/
func (tree *BinarySearchTree) Deep() int {
	leftDeep := 0
	rightDeep := 0
	/*if(tree.leftChild == nil && tree.rightChild == nil){
		return -1
	}*/
	if tree.leftChild.value == 0 && tree.rightChild.value == 0 {
		// 空树返回-1
		return -1
	}
	current := tree
	for current.leftChild.value != 0 {
		current = current.leftChild
		leftDeep++
	}
	current = tree
	for current.rightChild.value != 0 {
		current = current.rightChild
		rightDeep++
	}
	if current.leftChild.value != 0 {
		rightDeep++
	}
	if rightDeep > leftDeep {
		return rightDeep
	}
	return leftDeep
}

func (tree *BinarySearchTree) Height() int {

	return tree.Deep() + 1
}

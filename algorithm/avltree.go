package algorithm

// avl树  自平衡的二叉查找树

type AvlNode struct {
	value int
	left *AvlNode
	right *AvlNode
}

type AvlTree struct {
	root *AvlNode
}
/**
	构造一个avl树
 */
func CreateAvlTree() *AvlTree{

	return &AvlTree{}
}
/**
	添加一个节点
 */
func (avl *AvlTree) InsertAvlNode(value int){
	newNode := &AvlNode{
		value:value,
		left:nil,
		right:nil,
	}
	if avl.root == nil{
		avl.root =newNode
		return
	}
	tempNode := avl.root
	for{
		if tempNode .value > value{
			if tempNode.left == nil{
				tempNode.left = newNode
				break
			}
			tempNode = tempNode.left
		}
		if tempNode.value < value{
			if tempNode.right == nil{
				tempNode.right = newNode
				break
			}
			tempNode = tempNode.right
		}
	}
	// 判断是否旋转树


}

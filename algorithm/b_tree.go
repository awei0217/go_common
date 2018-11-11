package algorithm

/**
一、B树的定义
B树也称B-树,它是一颗多路平衡查找树。我们描述一颗B树时需要指定它的阶数，阶数表示了一个结点最多有多少个孩子结点，一般用字母m表示阶数。当m取2时，就是我们常见的二叉搜索树。
一颗m阶的B树定义如下：
1、每个结点最多有m-1个关键字。
2、根结点最少可以只有1个关键字。
3、非根结点至少有Math.ceil(m/2)-1个关键字。
4、每个结点中的关键字都按照从小到大的顺序排列，每个关键字的左子树中的所有关键字都小于它，而右子树中的所有关键字都大于它。
5、所有叶子结点都位于同一层，或者说根结点到每个叶子结点的长度都相同
*/
type BT struct {
	m      int   //树的度
	parent *BT   //指向父节点的指针
	keyNum int   //每个节点关键字的个数关键字个数
	key    []int //关键字向量
	child  []*BT //子树指针向量
}


/**
m 为树的度数
*/
func NewBTree(m int) *BT {
	return &BT{
		m:		m, //b树的度
		parent: nil,
		keyNum: 0,
		key:    make([]int,m-1,m-1),//每个节点的关键字个数  m/2-1 <= l <= m-1
		child:  make([]*BT, m, m), //每个节点子节点个数    m/2   <= l <= m
	}
}

func (bt *BT) Insert(key int) bool {
	if bt == nil{
		return false
	}
	return true
}
func Split(bt *BT) {

}



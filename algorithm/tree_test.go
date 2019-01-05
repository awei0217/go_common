package algorithm

import "testing"

func TestMinDepth(t *testing.T) {
	left := &TreeNode{1, nil, nil}
	right := &TreeNode{3, nil, nil}
	root := &TreeNode{2, left, right}

	t.Log(MinDepth(root))
}

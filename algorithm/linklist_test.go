package algorithm

import "testing"

func TestLinkList_Add(t *testing.T) {
	linklist := &LinkList{
		nil, nil, 0,
	}
	linklist.Add(1, 0)
	linklist.Add(2, 0)
	linklist.Add(3, 0)
	linklist.Add(4, 2)
	linklist.Add(4, 0)

	t.Log(linklist.Get(0))
	t.Log(linklist.Get(1))
	t.Log(linklist.Get(2))
	t.Log(linklist.Get(3))
	t.Log(linklist.Get(4))
	t.Log("值为1的下标是:", linklist.IndexOfFirst(1))
	t.Log("值为4的下标是:", linklist.IndexOfFirst(4))
	t.Log(linklist.RemoveFirstPop())
	t.Log(linklist.Get(0))
	t.Log(linklist.Get(1))
	t.Log(linklist.Get(2))
	t.Log(linklist.Get(3))
}

func TestLinkList_Reverse(t *testing.T) {
	list := &LinkList{}
	list.AddToFirst(1)
	list.AddToFirst(2)
	list.AddToFirst(3)
	list.Reverse()
	list.Print()
}

package algorithm

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	linkList := AddTail(nil, 1)
	linkList = AddTail(linkList, 2)
	linkList = AddTail(linkList, 3)
	linkList = AddTail(linkList, 4)
	linkList = AddTail(linkList, 5)

}

func TestRevers(t *testing.T) {
	linkList := AddTail(nil, 3)
	linkList = AddTail(linkList, 4)
	linkList = AddTail(linkList, 5)
	linkList = Revers(linkList)
	Print(linkList)
}

func TestDeleteCurrentNode(t *testing.T) {
	linkList := AddTail(nil, 3)
	linkList = AddTail(linkList, 4)
	linkList = DeleteCurrentNode(linkList)
	Print(linkList)
}

func TestFindCenterNode(t *testing.T) {
	linkList := AddTail(nil, 3)
	linkList = AddTail(linkList, 4)
	linkList = AddTail(linkList, 5)
	temp := FindCenterNode(linkList)
	fmt.Println(temp.Value)
}

func TestFindReciprocalNode(t *testing.T) {
	linkList := AddTail(nil, 3)
	linkList = AddTail(linkList, 4)
	linkList = AddTail(linkList, 5)
	temp := FindReciprocalNode(linkList, 1)
	if temp != nil {
		fmt.Println(temp.Value)
	}
	Print(linkList)
}

func TestDeleteNthFromEnd(t *testing.T) {
	linkList := AddTail(nil, 3)
	linkList = AddTail(linkList, 4)
	linkList = AddTail(linkList, 5)
	linkList = DeleteNthFromEnd(linkList, 3)
	Print(linkList)
}

func TestMergeSortLinkList(t *testing.T) {
	linklist1 := AddTail(nil, 1)
	linklist1 = AddTail(linklist1, 2)
	linklist1 = AddTail(linklist1, 4)
	linklist1 = AddTail(linklist1, 5)
	linklist2 := AddTail(nil, 0)

	three := MergeSortLinkList(nil, linklist2)
	Print(three)
}

func TestDeleteNode(t *testing.T) {
	linklist1 := AddTail(nil, 1)
	linklist1 = AddTail(linklist1, 2)
	linklist1 = AddTail(linklist1, 4)
	linklist1 = AddTail(linklist1, 1)
	linklist1 = DeleteAppointValueNode(linklist1, 1)
	Print(linklist1)
}

func TestSortOneLinkList(t *testing.T) {
	linklist1 := AddTail(nil, 1)
	linklist1 = AddTail(linklist1, 2)
	linklist1 = AddTail(linklist1, 4)
	linklist1 = AddTail(linklist1, 5)
	linklist1 = AddTail(linklist1, 0)
	Print(MergerSortOneLinkList(linklist1))
}

func TestInsertSortOneLinkList(t *testing.T) {
	linklist1 := AddTail(nil, 4)
	linklist1 = AddTail(linklist1, 2)
	linklist1 = AddTail(linklist1, 1)
	linklist1 = AddTail(linklist1, 3)
	Print(InsertSortOneLinkList(linklist1))

}

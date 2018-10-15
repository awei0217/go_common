package algorithm

import "fmt"

type OneLinkedList *OneNode

// 单链表
type OneNode struct {
	Next *OneNode
	Value int
}


/**
	单链表添加尾部
 */
func AddTail(linkList OneLinkedList, value int)OneLinkedList  {
	if linkList == nil{
		linkList = &OneNode{
			Next:nil,
			Value:value,
		}
		return linkList
	}
	head := linkList
	for linkList.Next != nil{
		linkList = linkList.Next
	}
	linkList.Next = &OneNode{
		Next:nil,
		Value:value,
	}
	return head
}
/**
	单链表添加头部
 */
func AddHead(linkList OneLinkedList,value int)OneLinkedList{
	if linkList == nil{
		linkList = &OneNode{Next:nil,Value:value}
		return linkList
	}
	node := &OneNode{Next:nil,Value:value}
	node.Next = linkList
	return node

}
/**
	& 是取地址符号 , 即取得某个变量的地址 , 如 ; &a
	* 是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值
	反转单链表
 */
func Revers(linkList OneLinkedList)OneLinkedList  {
	if linkList == nil{
		return linkList
	}
	var newNode =  &OneNode{
		Next:nil,
		Value:linkList.Value,
	}
	for linkList.Next != nil {
		next := *linkList.Next
		next.Next = newNode
		newNode = &next
		linkList = linkList.Next
	}
	return newNode
}
/**
	寻找中间节点
	//两个下标开始往后遍历，第一个步长为1，第二个步长为2，当第二个到最后的时候，第一个恰好到中间
 */
func FindCenterNode(linkList OneLinkedList) OneLinkedList {
	one := linkList
	two := linkList.Next
	for two != nil && two.Next != nil{
		one= one.Next
		two = two.Next.Next
	}
	return one
}
/**
	查找单链表倒数第n个节点
 */
func FindReciprocalNode(linkList OneLinkedList,n int)OneLinkedList{
    i:=0
    one := linkList
    two := linkList
    for two != nil{
		two = two.Next
		i++
		if i>n {
			one  = one.Next
		}
	}
    //查找不存在的节点
    if i< n{
    	return nil
	}
	return one
}
/**
	删除当前节点
 */
func DeleteCurrentNode(linkList OneLinkedList)OneLinkedList{
	next := linkList.Next
	if next == nil{
		linkList = next
		return linkList
	}
	linkList.Value = next.Value
	linkList.Next = next.Next
	return linkList
}
/**
	删除倒数第n个节点，返回头节点
 */
func DeleteNthFromEnd(linkList OneLinkedList, n int) OneLinkedList {
	i:=0
	one := linkList
	two := linkList
	for one.Next != nil{
		one = one.Next
		i++
		if i>n{
			two = two.Next
		}
	}
	// 删除头节点
	if n-i == 1{
		linkList = linkList.Next
		return linkList
	}
	// n 大于链表的长度，无法删除，返回整个链表
	if n -i > 1{
		return linkList
	}
	next := two.Next
	two.Next = next.Next
	return linkList
}
/**
	合并两个有序的链表
 */
func MergeSortLinkList(one,two OneLinkedList)OneLinkedList  {
	var head OneLinkedList
	var newNode OneLinkedList
	for one != nil && two != nil{
		if one.Value < two.Value{
			if head == nil{
				head = &OneNode{
					Value:one.Value,
					Next:nil,
				}
				newNode = head
			}else{
				head.Next = one
				head = head.Next
			}
			one = one.Next
		}else if one.Value > two.Value{
			if head == nil{
				head = &OneNode{
					Value: two.Value,
					Next:  nil,
				}
				newNode = head
			}else{
				head.Next = two
				head = head.Next
			}
			two = two.Next
		}else{
			if head == nil{
				head = &OneNode{
					Value: one.Value,
					Next:  nil,
				}
				newNode = head
				head.Next = two
				head = head.Next
				one = one.Next
				two = two.Next
			}else{
				head.Next = one
				one = one.Next
				head = head.Next
				head.Next = two
				two = two.Next
				head = head.Next
			}
		}
	}
	for one != nil{
		if head == nil{
			head = &OneNode{
				Next:nil,
				Value:one.Value,
			}
			newNode = head
		}else{
			head.Next  =one
			head =head.Next
		}
		one = one.Next
	}
	for two != nil{
		if head == nil {
			head = &OneNode{
				Next:  nil,
				Value: two.Value,
			}
			newNode = head
		}else{
			head.Next  =two
			head =head.Next
		}
		two = two.Next
	}
	return newNode
}
/**
	删除给定值的node
 */
func DeleteAppointValueNode(linklist OneLinkedList,value int)OneLinkedList  {
	if linklist == nil{
		return linklist
	}
	for linklist.Value == value{
		linklist = linklist.Next
		if linklist == nil{
			break
		}
	}
	temp := linklist
	for linklist != nil  {
		next := linklist.Next
		if next != nil{
			if linklist.Next.Value == value{
				linklist.Next = next.Next
				continue
			}
		}
		linklist = linklist.Next
	}
	return temp
}
/**
	用归并排序对单链表排序
 */
func MergerSortOneLinkList(linkList OneLinkedList)OneLinkedList  {
	if linkList == nil  || linkList.Next == nil{
		return linkList
	}
	first := linkList
	// 获取中间节点
	mid := FindCenterNode(linkList)
	second := mid.Next
	mid.Next = nil
	first = MergerSortOneLinkList(first)
	second = MergerSortOneLinkList(second)
	return MergeSortLinkList(first,second)
}
/**
	用插入排序对单链表进行排序
 */
func InsertSortOneLinkList(linklist OneLinkedList)OneLinkedList  {
	pre,cur :=linklist,linklist.Next
	for cur!= nil{
		for pre != cur {
			if cur.Value <= pre.Value{
				cur.Value,pre.Value = pre.Value,cur.Value
			}
			pre = pre.Next
		}
		cur = cur.Next
		pre = linklist
	}
	return linklist
}

/**
	打印单链表
 */
func Print(linkList OneLinkedList){
	for linkList != nil{
		fmt.Print(linkList.Value,",")
		linkList = linkList.Next
	}
	fmt.Println()
}

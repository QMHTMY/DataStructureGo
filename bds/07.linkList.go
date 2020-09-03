// go实现单链表
package bds // main

import (
	"fmt"
)

//节点结构： {data,next} data域存放节点值的数据域，next域存放节点的直接后继地址
type Node struct {
	Data interface{}
	Next *Node
}

func IsEmpty(node *Node) bool {
	return node == nil
}

func IsLast(node *Node) bool  {
	return node.Next == nil
}

func FindPrevious(data interface{}, node *Node) *Node  {
	tmp := node
	for tmp.Next != nil && tmp.Next.Data != data {
		tmp = tmp.Next
	}
	return tmp
}

// 查找值在哪个节点
func Find(data interface{}, node *Node) *Node  {
	tmp := node
	for tmp.Data != data {
		tmp = tmp.Next
	}
	return tmp
}

// 在指定节点插入节点
func Insert(data interface{}, position *Node)  {
	tmpCell := new(Node)
	if tmpCell == nil {
		fmt.Println("err: out of space")
	}

	tmpCell.Data = data
	tmpCell.Next = position.Next
	position.Next = tmpCell
}

func DeleteNode(data interface{}, node *Node)  {
	preview := FindPrevious(data, node)
	tmp := Find(data, node)
	preview.Next = tmp.Next
	tmp = nil
}

func DeleteList(node **Node)  {
	p := node
	for *p != nil {
		tmp := *p
		*p = nil
		*p = tmp.Next
	}
}

func PrintList(list *Node) {
	p := list
	for p != nil {
		fmt.Printf("%d-%v-%p ", p.Data, p, p.Next)
		p = p.Next
	}
	fmt.Println()
}

func main() {
    //测试时将bds改为main
	headNode := &Node{
		Data:0,
		Next:nil,
	}

	secondNode := &Node{
		Data:1,
		Next:nil,
	}

	thirdNode := &Node{
		Data:2,
		Next:nil,
	}

	list := headNode
	Insert(1, headNode)
	Insert(1, secondNode)
	Insert(1, thirdNode)

	PrintList(list)
	fmt.Println(IsEmpty(list))
	fmt.Println(IsLast(headNode))

	p := Find(0, list)
	Insert(2, p)
	PrintList(list)

	DeleteNode(1, list)
	PrintList(list)

	DeleteList(&list)
	PrintList(list)
}

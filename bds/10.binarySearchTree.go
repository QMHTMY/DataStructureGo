// go实现二叉搜索树

package bds //main

import (
	"fmt"
)

type Item interface {
}

// 节点
type Node struct {
	Key int
	Value Item
	left *Node
	right *Node
}

// 实现方法
type Noder interface {
	Insert(key int, value Item)
	Min() *Item
	Max() *Item
	Search(key int) bool
	PreOrderTraverse(f func(item Item))
	InOrderTraverse(f func(item Item))
	PostOrderTraverse(f func(item Item))
	String()
}

// 树结构
type ItemBinarySearchTree struct {
	Root *Node
}

// 各方法实现
func (bst *ItemBinarySearchTree) Insert(key int, value Item)  {
	n := &Node{key, value, nil, nil}

	if bst.Root == nil{
		bst.Root = n
	}else {
		insertNode(bst.Root, n)
	}
}

func insertNode(node, newNode *Node)  {
	if newNode.Key < node.Key{ // 左插入
		if node.left == nil{
			node.left = newNode
		}else {
			insertNode(node.left, newNode)
		}
	}else {                    // 右插入
		if node.right == nil{
			node.right = newNode
		}else {
			insertNode(node.right, newNode)
		}
	}
}
func (bst *ItemBinarySearchTree) Min() *Item {
	n := bst.Root

	if n == nil{
		return nil
	}

	for {
		if n.left == nil{
			return &n.Value
		}
		n = n.left
	}
}

// 最大值，一定在树的最右边
func (bst *ItemBinarySearchTree) Max() *Item  {
	n := bst.Root

	if n == nil{
		return &n.Value
	}

	for {
		if n.right == nil{
			return &n.Value
		}else {
			n = n.right
		}
	}
}

// 二叉搜索树，查找元素
func (bst *ItemBinarySearchTree) Search(key int) (bool)  {
	return search(bst.Root, key)
}

func search(n *Node, key int) bool  {
	if n == nil{
		return false
	}

	if key < n.Key{
		return search(n.left, key)
	}

	if key > n.Key{
		return search(n.right, key)
	}

	return true
}

// 前序遍历
func (bst *ItemBinarySearchTree) PreOrderTraverse(f func(item Item)) {
	preOrdertraver(bst.Root, f)
}

func preOrdertraver(n *Node, f func(item Item))  {
	if n != nil{
		f(n.Value)
		preOrdertraver(n.left, f)
		preOrdertraver(n.right, f)
	}
}

// 中序遍历
func (bst *ItemBinarySearchTree) InOrderTraverse(f func(item Item))  {
	inOrderTraver(bst.Root, f)
}

func inOrderTraver(n *Node, f func(item Item))  {
	if n != nil{
		inOrderTraver(n.left, f)
		f(n.Value)
		inOrderTraver(n.right, f)
	}
}

// 后序遍历
func (bst *ItemBinarySearchTree) PostOrderTraverse(f func(item Item))  {
	postOrderTraver(bst.Root, f)
}

func postOrderTraver(n *Node, f func(item Item))  {
	if n != nil{
		postOrderTraver(n.left, f)
		postOrderTraver(n.right, f)
		f(n.Value)
	}
}


// 字符串化
func (bst *ItemBinarySearchTree) String()  {
	fmt.Println("Bst--Bst--Bst--Bst--Bst--Bst--Bst--Bst--Bst--Bst")
	stringify(bst.Root, 0)
	fmt.Println("Bst--Bst--Bst--Bst--Bst--Bst--Bst--Bst--Bst--Bst")
}

func stringify(n *Node, level int)  {
	if n != nil{
		format := ""
		for i:=0; i<level; i++{
			format += "         "
		}

		format += "---["
		level ++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.Key)
		stringify(n.right, level)
	}
}

func fillTree(bst *ItemBinarySearchTree)  {
	bst.Insert(9, "9")
	bst.Insert(2, "2")
	bst.Insert(1, "1")
	bst.Insert(8, "8")
	bst.Insert(6, "6")
	bst.Insert(4, "4")
	bst.Insert(7, "7")
	bst.Insert(13, "13")
	bst.Insert(14, "14")
}

func main() {
    var bst ItemBinarySearchTree
	var result []string

	fillTree(&bst)
	bst.String()

	bst.PreOrderTraverse(func(i Item) {
		result = append(result, fmt.Sprintf("%s", i))
	})

    for _, item := range result {
        fmt.Println(item)
    }
}

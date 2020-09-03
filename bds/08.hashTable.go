// go实现简单哈希表
//package bds // main
package main

import (
    "fmt"
)

type Dict struct {
   Key   string
   Value interface{}
}

type MapNode struct {
   Data Dict
   Next *MapNode
}

// 定义数组链表(哈希表)
var Arr [16]*MapNode

// 初始化链表头节点
func NewNodeHead() *MapNode {
   node := new(MapNode)
   node.Data.Key = "head"
   node.Data.Value = "headvalue"
   node.Next = nil
   return node
}

// 初始化哈希表
func NewHash() {
   for i := 0; i < 16; i++ {
      Arr[i] = NewNodeHead()
   }
}

// 拉链法保存key-value
func (node *MapNode) data(k string, v interface{}) *MapNode {
   if node == nil {
      node = NewNodeHead()
   }
   node.Data.Key = k
   node.Data.Value = v
   return node
}

// 根据key取value
func (node *MapNode) getKey(k string) interface{} {
   if node.Next == nil {
      return nil
   }

   for node.Next != nil {
      if node.Next.Data.Key == k {
         return node.Next.Data.Value
      } else {
         node = node.Next
      }
   }

   return nil
}

// 添加新数据
func (node *MapNode) add(k string, v interface{}) {
   if node.getKey(k) != nil {       // 判断k是否已有
      return
   }

   for node.Next != nil {           //遍历到尾节点
      node = node.Next
   }

   node.Next = node.Next.data(k, v) // 创建新的尾节点
}

// 递归打印
func (node *MapNode) Log() {
   if node.Next == nil {
      return
   }
   fmt.Println(node.Next.Data)
   node.Next.Log()
}

// 计算链表长度
func (node *MapNode) Length() int {
   if node == nil {
      return 0
   }

   i := 0
   for node.Next != nil {
      i++
      node = node.Next
   }

   return i
}

func SetKey(k string, v interface{}) {
   num := hashNum(k)
   Arr[num].add(k, v)
}

func GetKey(k string) interface{} {
   num := hashNum(k)
   return Arr[num].getKey(k)
}

func hashNum(key string) int {
   var index int = 0
   index = int(key[0])

   // 累积新的数值
   for k := 0; k < len(key); k++ {
      index *= (1103515245 + int(key[k]))
   }

   index >>= 27 // 右位移27位，可位移任何数，只需别太大导致清空

   // 按位与，这里要求32以内的数，则32-1即可，也就是说
   // index跟1111进行&，得到15以内的数
   // index跟11111取余，则得31以内的数
   index &= 16 - 1

   return index
}

// 测试
func main()  {
   node := NewNodeHead()
   node.add("1","2")
   node.add("2", 3)
   node.Log()
   fmt.Println(node.Length())
}

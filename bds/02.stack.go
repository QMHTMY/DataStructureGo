//go栈实现

package bds //main

import (
    "fmt"
)

//基本数据容器
type Item interface {}

type Stack struct {
    items []Item
}

//栈操作函数
func (s *Stack) New() (*Stack) {
    s.items = []Item{}
    return s
}

func (s *Stack) Push(item Item) {
    s.items = append(s.items, item)
}

func (s *Stack) Pop() (Item) {
    res := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return res
}

func (s *Stack) Peek() (Item) {
    data := s.items[len(s.items)-1]
    return data
}

func (s *Stack) Size() (int) {
    return len(s.items)
}

func (s *Stack) IsEmpty() (bool) {
    return len(s.items) == 0
}

// 初始化栈
var stack Stack

func initStack() *Stack {
    if stack.items == nil {
        stack = Stack{}
        stack.New()
    }
    return &stack
}

// 测试
func main() {
    //测试此函数时修改bds为main
    stk := initStack()
    stk.Push(1)
    stk.Push(5)
    stk.Push(6)
    fmt.Println("size ", stk.Size())
    fmt.Println("top ", stk.Peek())
    fmt.Println("rm top ", stk.Pop())
    fmt.Println("size ", stk.Size())
}

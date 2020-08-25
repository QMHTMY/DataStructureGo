//go队列实现

package bds //main

import (
    "fmt"
)

//基本数据容器
type Item interface {}

type Queue struct {
    items []Item
}

//队列操作函数
func (q *Queue) New() (*Queue) {
    q.items = []Item{}
    return q
}

func (q *Queue) Push(item Item) {
    q.items = append(q.items, item)
}

func (q *Queue) Pop() (Item) {
    res := q.items[0]
    q.items = q.items[1:]
    return res
}

func (q *Queue) Size() (int) {
    return len(q.items)
}

func (q *Queue) IsEmpty() (bool) {
    return len(q.items) == 0
}

// 初始化队列
var queue Queue

func initQueue() (*Queue) {
    if queue.items == nil {
        queue = Queue{}
        queue.New()
    }
    return &queue
}

// 测试
func main() {
    //测试此函数时修改bds为main
    q := initQueue()
    q.Push(1)
    q.Push(5)
    q.Push(6)
    fmt.Println("size ", q.Size())
    fmt.Println("rm head", q.Pop())
    fmt.Println("size ", q.Size())
}

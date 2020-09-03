//go双端队列实现

package bds //main

import (
    "fmt"
    "sync"
)

//基本数据容器
type Item interface {}

type Deque struct {
    items []Item
    sync.RWMutex
}

//双端队列操作函数
func (q *Deque) New() (*Deque) {
    q.items = []Item{}
    return q
}

func (q *Deque) addFront(item Item) {
    q.Lock()
    defer q.Unlock()
    q.items = append(q.items, item)
}

func (q *Deque) addRear(item Item){
    q.Lock()
    defer q.Unlock()
    temp := []Item{item}
    q.items = append(temp, q.items...)
}

func (q *Deque) removeFront() (Item) {
    q.Lock()
    defer q.Unlock()
    res := q.items[0]
    q.items = q.items[1:]
    return res
}

func(q *Deque) removeRear() (Item) {
    q.Lock()
    defer q.Unlock()
    res := q.items[len(q.items)-1]
    q.items = q.items[:len(q.items)-1]
    return res
}

func (q *Deque) Size() (int) {
    return len(q.items)
}

func (q *Deque) IsEmpty() (bool) {
    return len(q.items) == 0
}

// 初始化双端队列
var deque Deque

func initDeque() (*Deque) {
    if deque.items == nil {
        deque = Deque{}
        deque.New()
    }
    return &deque
}

// 测试
func main() {
    //测试此函数时修改bds为main
    dq := initDeque()

    dq.addRear(1)
    dq.addRear(10)
    dq.addFront(5)
    dq.addFront(6)

    fmt.Println("size ", dq.Size())
    fmt.Println("head ", dq.removeFront())
    fmt.Println("tail ", dq.removeRear())
    fmt.Println("size ", dq.Size())
}

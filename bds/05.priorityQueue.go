//go优先队列的实现

package bds //main 

import (
    "fmt"
    "container/heap"
)

type entry struct {
    key      string
    index    int
    priority int
}

type PQueue []*entry

func (pq *PQueue) Len() int {
    return len(pq)
}

func (pq *PQueue) Less(i, j int) (bool) {
    return pq[i].priority < pq[j].priority
}

func (pq *PQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PQueue) Push(item interface{}) {
    tmp := item.(*entry)
    tmp.index = len(*pq)
    *pq = append(*pq, tmp)
}

func (pq *PQueue) Pop() interface{} {
    tmp := (*pq)[len(*pq)-1]
    tmp.index = -1
    *pq = (*pq)[:len(*pq)-1]
    return tmp
}

func (pq *PQueue) update(item *entry, value string, priority int) {
    item.key = value
    item.priority = priority
    heap.Fix(pq, item.index)
}

// go实现整数集合set
// 由于泛型还不支持，所以集合只能按照数据类型实现
// 字符串，整数，浮点数要分别实现一套集合
// 2021年8月泛型支持后，就可以只用实现一套

package main

import (
    "fmt"
    "sync"
    "sort"
)

type Intset struct {
    m map[int]bool
    sync.RWMutex
}

func NewSet() (*Intset) {
    return &Intset{ m: map[int]bool{} }
}

func (s *Intset) Add(item int) {
    s.Lock()
    defer s.Unlock()
    s.m[item] = true
}

func (s *Intset) Remove(item int) {
    s.Lock()
    defer s.Unlock()
    delete(s.m, item)
}

func (s *Intset) Has(item int) (bool) {
    s.RLock()
    defer s.RUnlock()
    _, ok := s.m[item]
    return ok
}

func (s *Intset) Len() (int) {
    return len(s.List())
}

func (s *Intset) Clear() {
    s.Lock()
    defer s.Unlock()
    s.m = map[int]bool{}
}

func (s *Intset) IsEmpty() (bool) {
    if s.Len() == 0 {
        return true
    }
    return false
}

func (s *Intset) List() ([]int) {
    s.RLock()
    defer s.RUnlock()
    list := []int{}
    for item := range s.m {
        list = append(list, item)
    }
    return list
}

func (s *Intset) SortList() ([]int) {
    s.RLock()
    defer s.RUnlock()
    list := []int{}
    for item := range s.m {
        list = append(list, item)
    }
    sort.Ints(list)
    return list
}

func main() {
    s := NewSet()

    s.Add(1)
    s.Add(2)
    s.Add(3)
    s.Add(4)
    s.Add(1)

    if s.Has(2) {
        fmt.Println("2 is in set")
    }

    fmt.Println("无序：", s.List())
    fmt.Println("有序：", s.SortList())
}

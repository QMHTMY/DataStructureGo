// go实现字符串集合set
// 目前只是结合泛型草案实现的集合
// 2021年8月泛型支持后，此代码才可用，先写在这里

package main

import (
    "fmt"
    "sync"
    "sort"
)

type Strset struct {
    m map[type T]bool
    sync.RWMutex
}

func NewSet() (*Strset) {
    return &Strset{ m: map[T]bool{} }
}

func (s *Strset) Add(item T) {
    s.Lock()
    defer s.Unlock()
    s.m[item] = true
}

func (s *Strset) Remove(item T) {
    s.Lock()
    defer s.Unlock()
    delete(s.m, item)
}

func (s *Strset) Has(item T) (bool) {
    s.RLock()
    defer s.RUnlock()
    _, ok := s.m[item]
    return ok
}

func (s *Strset) Len() (int) {
    return len(s.List())
}

func (s *Strset) Clear() {
    s.Lock()
    defer s.Unlock()
    s.m = map[T]bool{}
}

func (s *Strset) IsEmpty() (bool) {
    if s.Len() == 0 {
        return true
    }
    return false
}

func (s *Strset) List() ([]T) {
    s.RLock()
    defer s.RUnlock()
    list := []T{}
    for item := range s.m {
        list = append(list, item)
    }
    return list
}

func (s *Strset) SortList() ([]T) {
    s.RLock()
    defer s.RUnlock()
    list := []T{}
    for item := range s.m {
        list = append(list, item)
    }
    sort.Strings(list)
    return list
}

func main() {
    s := NewSet()

    s.Add("b")
    s.Add("a")
    s.Add("c")
    s.Add("d")
    s.Add("a")

    if s.Has("a") {
        fmt.Println("a is in set")
    }

    fmt.Println("无序：", s.List())
    fmt.Println("有序：", s.SortList())
}

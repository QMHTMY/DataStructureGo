// go实现字符串集合set
// 由于泛型还不支持，所以集合只能按照数据类型实现
// 字符串，整数，浮点数要分别实现一套集合
// 2021年8月泛型支持后，就可以只用实现一套

package main

import (
    "fmt"
    "sync"
    "sort"
)

type Strset struct {
    m map[string]bool
    sync.RWMutex
}

func NewSet() (*Strset) {
    return &Strset{ m: map[string]bool{} }
}

func (s *Strset) Add(item string) {
    s.Lock()
    defer s.Unlock()
    s.m[item] = true
}

func (s *Strset) Remove(item string) {
    s.Lock()
    defer s.Unlock()
    delete(s.m, item)
}

func (s *Strset) Has(item string) (bool) {
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
    s.m = map[string]bool{}
}

func (s *Strset) IsEmpty() (bool) {
    if s.Len() == 0 {
        return true
    }
    return false
}

func (s *Strset) List() ([]string) {
    s.RLock()
    defer s.RUnlock()
    list := []string{}
    for item := range s.m {
        list = append(list, item)
    }
    return list
}

func (s *Strset) SortList() ([]string) {
    s.RLock()
    defer s.RUnlock()
    list := []string{}
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

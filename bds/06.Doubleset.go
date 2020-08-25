// go实现浮点数集合set
// 由于泛型还不支持，所以集合只能按照数据类型实现
// 字符串，整数，浮点数要分别实现一套集合
// 2021年8月泛型支持后，就可以只用实现一套

package bds //main

import (
    "fmt"
    "sync"
    "sort"
)

type Floatset struct {
    m map[float64]bool
    sync.RWMutex
}

func NewSet() (*Floatset) {
    return &Floatset{ m: map[float64]bool{} }
}

func (s *Floatset) Add(item float64) {
    s.Lock()
    defer s.Unlock()
    s.m[item] = true
}

func (s *Floatset) Remove(item float64) {
    s.Lock()
    defer s.Unlock()
    delete(s.m, item)
}

func (s *Floatset) Has(item float64) (bool) {
    s.RLock()
    defer s.RUnlock()
    _, ok := s.m[item]
    return ok
}

func (s *Floatset) Len() (int) {
    return len(s.List())
}

func (s *Floatset) Clear() {
    s.Lock()
    defer s.Unlock()
    s.m = map[float64]bool{}
}

func (s *Floatset) IsEmpty() (bool) {
    if s.Len() == 0 {
        return true
    }
    return false
}

func (s *Floatset) List() ([]float64) {
    s.RLock()
    defer s.RUnlock()
    list := []float64{}
    for item := range s.m {
        list = append(list, item)
    }
    return list
}

func (s *Floatset) SortList() ([]float64) {
    s.RLock()
    defer s.RUnlock()
    list := []float64{}
    for item := range s.m {
        list = append(list, item)
    }
    sort.Float64s(list)
    return list
}

func main() {
    //测试本代码时修改bds为main
    s := NewSet()

    s.Add(1.2)
    s.Add(2.2)
    s.Add(3.2)
    s.Add(4.2)
    s.Add(1.2)

    if s.Has(2.2) {
        fmt.Println("2.2 is in set")
    }

    fmt.Println("无序：", s.List())
    fmt.Println("有序：", s.SortList())
}

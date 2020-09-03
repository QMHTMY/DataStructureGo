//go小顶堆实现

package bds //main 

import (
    "fmt"
)

type Heap struct {
    list []int
    size int
}

type Method interface {
    BuildHeap(alist []int)
    Insert(k int)
    PercUp(i int)
    PercDown(i int)
    MinChild(i int) (int)
    DelMin() (int)
    IsEmpty() (bool)
}

func (hp *Heap) BuildHeap(alist []int) {
    hp.list = append(hp.list, alist...)
    hp.size = len(alist)
    i := hp.size / 2

    for {
        if i <= 0 {
            break
        }

        hp.PercDown(i)
        i -= 1
    }
}

func (hp *Heap) PercDown(i int) {
    for {
        if (i * 2) > hp.size {
            break
        }

        mc := hp.MinChild(i)
        if hp.list[i] > hp.list[mc] {
            tmp := hp.list[i]
            hp.list[i] = hp.list[mc]
            hp.list[mc] = tmp
        }
        i = mc
    }
}

func (hp *Heap) MinChild(i int) (int) {
    if i * 2 + 1 > hp.size {
        return i*2
    } else {
        if hp.list[i*2] < hp.list[i*2 + 1] {
            return i*2
        } else {
            return i*2+1
        }
    }
}

func (hp *Heap) PercUp(i int) {
    for {

        if i / 2 <= 0 {
            break
        }

        if hp.list[i] < hp.list[i/2] {
            tmp := hp.list[i/2]
            hp.list[i/2] = hp.list[i]
            hp.list[i] = tmp
        }

        i /= 2
    }
}

func (hp *Heap) Insert(k int) {
    hp.list = append(hp.list, k)
    hp.size += 1
    hp.PercUp(hp.size)
}

func (hp *Heap) DelMin() (int) {
    retval := hp.list[1]

    hp.list[1] = hp.list[hp.size]
    hp.list = hp.list[:hp.size]
    hp.size -= 1
    hp.PercDown(1)

    return retval
}

func (hp *Heap) IsEmpty() (bool) {
    return hp.size == 0
}

func main() {
    //测试时，修改bds为main
    hp := Heap{list: []int{0}, size: 0}
    fmt.Println(hp.size)

    alist := []int{1,7,2,6,3,4,5,8}
    hp.BuildHeap(alist)
    fmt.Println(hp.size)

    fmt.Println(hp.DelMin())
    fmt.Println(hp.IsEmpty())
}

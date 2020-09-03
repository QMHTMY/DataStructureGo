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
    buildHeap(alist []int)
    insert(k int)
    percUp(i int)
    percDown(i int)
    minChild(i int) (int)
    delMin() (int)
    isEmpty() (bool)
}

func (hp *Heap) buildHeap(alist []int) {
    hp.list = append(hp.list, alist...)
    hp.size = len(alist)
    i := hp.size / 2

    for {
        if i <= 0 {
            break
        }

        hp.percDown(i)
        i -= 1
    }
}

func (hp *Heap) percDown(i int) {
    for {
        if (i * 2) > hp.size {
            break
        }

        mc := hp.minChild(i)
        if hp.list[i] > hp.list[mc] {
            tmp := hp.list[i]
            hp.list[i] = hp.list[mc]
            hp.list[mc] = tmp
        }
        i = mc
    }
}

func (hp *Heap) minChild(i int) (int) {
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

func (hp *Heap) percUp(i int) {
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

func (hp *Heap) insert(k int) {
    hp.list = append(hp.list, k)
    hp.size += 1
    hp.percUp(hp.size)
}

func (hp *Heap) delMin() (int) {
    retval := hp.list[1]

    hp.list[1] = hp.list[hp.size]
    hp.list = hp.list[:hp.size]
    hp.size -= 1
    hp.percDown(1)

    return retval
}

func (hp *Heap) isEmpty() (bool) {
    return hp.size == 0
}

func main() {
    //测试时，修改bds为main
    hp := Heap{list: []int{0}, size: 0}
    fmt.Println(hp.size)

    alist := []int{1,7,2,6,3,4,5,8}
    hp.buildHeap(alist)
    fmt.Println(hp.size)

    fmt.Println(hp.delMin())
    fmt.Println(hp.isEmpty())
}

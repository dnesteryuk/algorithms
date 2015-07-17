package main

import "fmt"

type heap struct {
    list [13]int
    N int
}

func (h *heap) insert(val int) {
    h.list[h.N] = val

    h.swim(h.N)
    h.N = h.N + 1
}

func (h *heap) swim(child int) {
    parent := (child - 1) / 2

    fmt.Println(h.list[parent], h.list[child])
    fmt.Println(h.list[parent], h.list[child])

    if h.list[parent] < h.list[child] {
        h.swap(parent, child)

        fmt.Println(h.list)

        h.swim(parent)
    }
}

func (h *heap) swap(a, b int) {
    temp := h.list[a]
    h.list[a] = h.list[b]
    h.list[b] = temp
}

func (h *heap) deleteMax() {
    h.N = h.N - 1

    h.swap(0, h.N)

    h.list[h.N] = 0

    h.sink(0)
}

func (h *heap) sink(rootIndex int) {
    leftIndex := rootIndex * 2 + 1

    if leftIndex < h.N {
        rightIndex := rootIndex * 2 + 2

        root := h.list[rootIndex]
        left := h.list[leftIndex]

        var right int

        if rightIndex < h.N {
            right = h.list[rightIndex]
        } else {
            right = 0
        }

        if root < left && root < right {
            if left < right {
                h.swap(rootIndex, rightIndex)
                h.sink(rightIndex)
            } else {
                h.swap(rootIndex, leftIndex)
                h.sink(leftIndex)
            }
        } else if root < left {
            h.swap(rootIndex, leftIndex)
            h.sink(leftIndex)
        } else {
            h.swap(rootIndex, rightIndex)
            h.sink(rightIndex)
        }
    }
}

func main() {
    arr := [13]int{89, 80, 69, 73, 44, 32, 43, 16, 62, 18}

    h := heap{list: arr, N: 10}

    // 77 52 42

    // h.insert(77)
    // h.insert(52)
    // h.insert(42)

    fmt.Println("changed list ", h.list, len(h.list))
}
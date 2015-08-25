// This data structure is based on the binary heap
// https://en.wikipedia.org/wiki/Binary_heap
package dstructures

type heap struct {
    list []int
    N int
}

func (h *heap) Insert(val int) {
    h.list[h.N] = val

    h.Swim(h.N)
    h.N = h.N + 1
}

func (h *heap) Swim(child int) {
    parent := (child - 1) / 2

    if h.list[parent] < h.list[child] {
        h.Swap(parent, child)
        h.Swim(parent)
    }
}

func (h *heap) Swap(a, b int) {
    temp := h.list[a]
    h.list[a] = h.list[b]
    h.list[b] = temp
}

func (h *heap) deleteMax() {
    h.N = h.N - 1

    h.Swap(0, h.N)

    h.list[h.N] = 0

    h.sink(0)
}

func (h heap) sink(rootIndex int) {
    childIndex := rootIndex * 2 + 1

    for childIndex <= h.N {
        if childIndex < h.N && h.list[childIndex] < h.list[childIndex + 1] {
            childIndex = childIndex + 1 // right node is bigger than the left one, we are interested only in this one
        }

        if h.list[childIndex] < h.list[rootIndex] {
            break
        }

        h.Swap(rootIndex, childIndex)

        childIndex = childIndex + 1
    }
}

func BMaxHeap(list [] int, n int) heap {
    return heap{list: list, N: n}
}

// func main() {
//     arr := [13]int{89, 80, 69, 73, 44, 32, 43, 16, 62, 18}

//     h := heap{list: arr[:], N: 10}

//     h.Insert(77)
//     h.Insert(52)
//     h.Insert(42)

//     fmt.Println("new list ", h.list)
// }
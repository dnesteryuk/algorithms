// https://en.wikipedia.org/wiki/Merge_sort

package main

import "fmt"

func sort(list []int) {
    length := len(list)
    cache  := make([]int, length)

    mergeSort(list, cache, 0, length - 1)
}

func mergeSort(list []int, cache []int, lt, gt int) {
    if lt >= gt {
        return
    }

    mid := lt + (gt - lt) / 2

    /*
    Splits the list on the left and right parts
    to sort them separately.
    */
    mergeSort(list, cache, lt, mid)
    mergeSort(list, cache, mid + 1, gt)

    /*
    Merges the left and right part together.
    The left and right parts are already sorted.
    But, we need to merge them properly.
    Example of the list on this step:
      [4,5,9, 1,6,7]
    */
    merge(list, cache, lt, mid, gt)
}

func merge(list []int, cache []int, lt, mid, gt int) {
    /*
    Since the original list is changed during the merge
    we have to copy items to the temporary list.
    */
    for i := lt; i <= gt; i++ {
        cache[i] = list[i]
    }

    i := lt
    j := mid + 1

    for k := lt; k <= gt; k++ {
        /*
        The left part is exhausted,
        hence we just copy the right part.
        */
        if i > mid {
            list[k] = cache[j]
            j++

        /*
        The right part is exhausted,
        hence we just copy the left part.
        */
        } else if j > gt {
            list[k] = cache[i]
            i++

        /*
        Compares items from both sides
        makes a decision which item must be copied
        to the resulting list.
        */
        } else if cache[i] > cache[j] {
            list[k] = cache[j]
            j++
        } else {
            list[k] = cache[i]
            i++
        }
    }
}

func main() {
    l := []int{9, 2, 5, 4, 7, 6}

    sort(l)

    fmt.Println("Sorted list ", l)
}
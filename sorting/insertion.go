// https://en.wikipedia.org/wiki/Insertion_sort
package main

import "fmt"

func sort(list []int) {
    length := len(list)

    for i := 1; i < length; i++ {
        // Once we found that the previous node is
        // greater then the current one, we need to
        // swap elements until the left part of the array
        // does not become sorted.
        if list[i - 1] > list[i] {
            for j := i; j > 0; j-- {
                if list[j - 1] > list[j] {
                    swap(list, j - 1, j);
                } else {
                    break
                }
            }
        }
    }
}

func swap(list []int, i int, j int) {
    tmp := list[i]
    list[i] = list[j]
    list[j] = tmp
}

func main() {
    l := []int{9, 2, 5, 4, 3, 1, 8, 7, 0, 6}

    sort(l)

    fmt.Println("Sorted list ", l)
}
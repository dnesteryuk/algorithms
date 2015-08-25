// https://en.wikipedia.org/wiki/Selection_sort
package main

import "fmt"

func sort(list []int) {
    length := len(list)

    for i := 0; i < length; i++ {
        min := i

        for j := i + 1; j < length; j++ {
            if list[min] > list[j] {
                min = j
            }
        }

        if i != min {
            swap(list, i, min);
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
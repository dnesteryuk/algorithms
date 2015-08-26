// https://en.wikipedia.org/wiki/Shellsort
package main

import "fmt"

func sort(list []int) {
    length := len(list)
    steps  := steps(length)

    for k := len(steps) - 1; k >= 0; k-- {
        step := steps[k]

        for i := step; i < length; i++ {
            if list[i - step] > list[i] {
                for j := i; j >= step; j -= step {
                    if list[j - step] > list[j] {
                        swap(list, j - 1, j);
                    } else {
                        break
                    }
                }
            }
        }
    }
}

func steps(n int) []int {
    l := []int{}

    for i := 0; i < n / 3; {
        i = i * 3 + 1

        l = append(l, i)
    }

    return l
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
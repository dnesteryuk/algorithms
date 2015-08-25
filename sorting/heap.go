package main

import(
  "fmt"
  "dstructures"
)

func main() {
    arr := [13]int{89, 80, 69, 73, 44, 32, 43, 16, 62, 18}

    h := dstructures.BMaxHeap(arr[:], 10)

    h.Insert(77)
    h.Insert(52)
    h.Insert(42)

    fmt.Println("new list ", arr)
}
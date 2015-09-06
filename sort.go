package main

import(
    "flag"
    "fmt"
    "sorting"
    "cli"
)

func main() {
    alg := *flag.String("algorithm", "merge", "The algorithm to apply for sorting")

    flag.Parse()

    cli.RequireSortingArgs()
    list := cli.IntsArgs()

    if alg == "merge" {
        sorting.MergeSort(list)
    } else if alg == "selection" {
        sorting.Selection(list)
    } else if alg == "insertion" {
        sorting.Insertion(list)
    }

    fmt.Println("Algorithm:", alg)
    fmt.Println("Sorted list:", list)
}

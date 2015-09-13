package main

import(
    "flag"
    "fmt"
    "sorting"
    "cli"
)

func sort(list []int, f func([]int)) {
    f(list)
}

func main() {
    pAlg := flag.String("a", "merge", "The algorithm to apply for sorting [selection, insertion, shell, merge]")


    flag.Parse()

    cli.RequireSortingArgs()
    list := cli.IntsArgs()

    alg := *pAlg

    m := map[string]func(list []int) {
        "selection": sorting.Selection,
        "insertion": sorting.Insertion,
        "shell": sorting.ShellSort,
        "merge": sorting.MergeSort,
    }

    sort(list, m[alg])

    fmt.Println("Algorithm:", alg)
    fmt.Println("Sorted list:", list)
}

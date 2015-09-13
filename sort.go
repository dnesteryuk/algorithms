package main

import(
    "fmt"
    "sorting"
    "cli"
)

var availableAlgs = map[string]func(a []int) {
    "selection": sorting.Selection,
    "insertion": sorting.Insertion,
    "shell": sorting.ShellSort,
    "merge": sorting.MergeSort,
}

func main() {
    alg := cli.AlgorithmFlag(availableAlgs, "merge")
    cli.RequireSortingArgs()

    list := cli.IntsArgs()

    f := availableAlgs[alg]

    f(list)

    fmt.Println("Algorithm:", alg)
    fmt.Println("Sorted list:", list)
}

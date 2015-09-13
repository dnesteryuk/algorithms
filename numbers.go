package main

import(
    "fmt"
    "numbers"
    "cli"
)

var availableAlgs = map[string]func(a []int) {
    "eratosphen_sieve": numbers.ESieve,
}

func main() {
    alg := cli.AlgorithmFlag(availableAlgs, "eratosphen_sieve")
    args := cli.IntsArgs()

    f := availableAlgs[alg]

    fmt.Println("Algorithm:", alg)
    f(args)
}

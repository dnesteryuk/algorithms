package cli

import(
    "flag"
    "os"
    "fmt"
    "strconv"
    "exts"
)

func RequireArgs(minCount int, msg string) {
    count := len(flag.Args())

    if count < minCount {
        fmt.Println(msg)
        os.Exit(3)
    }
}

func RequireSortingArgs() {
    RequireArgs(2, "You must provide at least 3 elements to sort them")
}

func IntsArgs() []int {
    args := flag.Args()
    count := len(args)

    list := make([]int, count)

    for i := 0; i < count; i++ {
        el, _ := strconv.Atoi(args[i])

        list[i] = el
    }

    return list
}

func AlgorithmFlag(algs map[string]func(a []int), defaultAlg string) string {
    possibleAlgs := exts.KeysAsString(algs)
    alg := flag.String("a", defaultAlg, "The algorithm to be applied [" + possibleAlgs + "]")
    flag.Parse()
    return *alg
}
package cli

import(
    "flag"
    "os"
    "fmt"
    "strconv"
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
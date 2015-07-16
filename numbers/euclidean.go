// https://en.wikipedia.org/wiki/Euclidean_algorithm

package main

import "os"
import "fmt"
import "strconv"

func main() {
    a, err := strconv.Atoi(os.Args[1])

    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }

    b, err := strconv.Atoi(os.Args[2])

    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }

    for {
        x := a % b

        if x == 0 {
            fmt.Println(b)
            break
        } else {
            a = b
            b = x
        }
    }
}
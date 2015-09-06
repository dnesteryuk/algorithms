// https://en.wikipedia.org/wiki/Euclidean_algorithm

package main

import(
    "fmt"
    "cli"
    "flag"
)

func Gcd(a, b int) int {
    for {
        x := a % b

        if x == 0 {
            return b
        } else {
            a = b
            b = x
        }
    }
}

func main() {
    flag.Parse()

    cli.RequireArgs(2, "You must provide 2 distinct numbers as arguments to the program")
    args := cli.IntsArgs()

    fmt.Println("Greatest common divisor is", Gcd(args[0], args[1]))
}
// https://en.wikipedia.org/wiki/Euclidean_algorithm

package main

import "os"
import "fmt"
import "strconv"

func gcd(a, b int) int {
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
    if len(os.Args) < 3 {
        fmt.Println("You must provide 2 distinct numbers as arguments to the program")
        os.Exit(3)
    }

    a, _ := strconv.Atoi(os.Args[1])
    b, _ := strconv.Atoi(os.Args[2])

    fmt.Println("Greatest common divisor is", gcd(a, b))
}
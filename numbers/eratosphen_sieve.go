// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes

package numbers

import(
    "fmt"
    "math"
    "cli"
)

func ESieve(args []int) {
    cli.RequireArgs(1, "The number representing the limit is required.")

    var start, end int

    if len(args) == 1 {
        start = 2
        end = args[0]

    } else {
        start = args[0]
        end = args[1]
    }

    if start > end {
        start, end = end, start
    }

    // One is not a prime number
    if start == 1 {
        start = 2
    }

    numbers := GenPrimeNumbers(start, end)
    fmt.Println("Prime numbers:", numbers)
}

func GenPrimeNumbers(start, end int) []int {
    list := make([]int, 0)

    for i := 2; i <= end; i++ {
        list = append(list, i)
    }

    max := int(math.Sqrt(float64(end)))

    for k := 2; k <= max; k++ {
        i := k - 2

        if list[i] != 0 {
            j := k * k - 2

            for j + 2 <= end {
                list[j] = 0

                j = j + k
            }
        }
    }

    numbers := make([]int, 0)

    for i := 0; i < len(list); i++ {
        if list[i] != 0 && list[i] >= start {
            numbers = append(numbers, list[i])
        }
    }

    return numbers
}
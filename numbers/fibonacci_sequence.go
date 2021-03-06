// https://en.wikipedia.org/wiki/Fibonacci_number
package numbers

func GenFibonacci(limit int) []int {
    list := make([]int, 0)

    f := fibonacci()

    for i := 1; i < limit; i++ {
        list = append(list, f())
    }

    return list
}

func fibonacci() func() int {
    a := 1
    b := 0

    return func() int {
        x := a + b

        a = b
        b = x

        return x
    }
}
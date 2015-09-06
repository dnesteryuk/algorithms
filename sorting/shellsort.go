// https://en.wikipedia.org/wiki/Shellsort
package sorting

func ShellSort(list []int) {
    length := len(list)
    steps  := steps(length)

    for k := len(steps) - 1; k >= 0; k-- {
        step := steps[k]

        for i := step; i < length; i++ {
            if list[i - step] > list[i] {
                for j := i; j >= step; j -= step {
                    if list[j - step] > list[j] {
                        swap(list, j - 1, j);
                    } else {
                        break
                    }
                }
            }
        }
    }
}

func steps(n int) []int {
    l := []int{}

    for i := 0; i < n / 3; {
        i = i * 3 + 1

        l = append(l, i)
    }

    return l
}
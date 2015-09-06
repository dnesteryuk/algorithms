// https://en.wikipedia.org/wiki/Insertion_sort
package sorting

func Insertion(list []int) {
    length := len(list)

    for i := 1; i < length; i++ {
        /*
        Once we find that the previous node is
        greater then the current one, we swaps
        elements of the left part until the
        left part of the array is sorted.
        */
        if list[i - 1] > list[i] {
            for j := i; j > 0; j-- {
                if list[j - 1] > list[j] {
                    swap(list, j - 1, j);
                } else {
                    break
                }
            }
        }
    }
}

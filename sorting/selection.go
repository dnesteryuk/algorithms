// https://en.wikipedia.org/wiki/Selection_sort
package sorting

func Selection(list []int) {
    length := len(list)

    for i := 0; i < length; i++ {
        min := i

        // Looks for the smallest item
        for j := i + 1; j < length; j++ {
            if list[min] > list[j] {
                min = j
            }
        }

        // Swaps the current element with the smallest item
        if i != min {
            swap(list, i, min);
        }
    }
}
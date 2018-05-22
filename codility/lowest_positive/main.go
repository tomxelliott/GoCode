// return lowest positive number greater than 0 not included in the array
package main

import "sort"

func Solution(A []int) int {
    result := 1
    sort.Sort(sort.IntSlice(A))
    for _, i := range A {
        if i == result {
            result++
        }
    }
    return result
}

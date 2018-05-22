// lowest value of odd occurence in array
package main

import "sort"

func Solution(A []int) int {
    helperMap := make(map[int]int)
    for _, i := range A {
        if _, ok := helperMap[i]; ok {
            helperMap[i]++
        } else {
            helperMap[i] = 1
        }
    }
    
    var odds []int
    for k, v := range helperMap {
        if v % 2 != 0 {
            odds = append(odds, k)
        } 
    }
    sort.Sort(sort.IntSlice(odds))
    return odds[0]
}

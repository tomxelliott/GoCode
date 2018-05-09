package binarygap

import "strconv"

func binaryGap(N int) int {
    var binaryGap int
    var largestBinaryGap int
    
    binaryN := strconv.FormatInt(int64(N), 2)
    
    for _, i := range binaryN {
        if i == '0' {
            binaryGap += 1
        } else {
            if binaryGap > largestBinaryGap {
                largestBinaryGap = binaryGap
            }
            binaryGap = 0
        }
    }
    
    return largestBinaryGap
}

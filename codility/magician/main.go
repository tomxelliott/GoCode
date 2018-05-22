package main


// solution from codility that shuffles cards
func Solution(A []int) []int {
    var result []int
    top := A[0:len(A)/2]
    bottom := A[len(A)/2:]
    
    for i := 0; i < len(A)/2; i++ {
        result = append(result, bottom[i])
        result = append(result, top[i])
    }
    return result    
}

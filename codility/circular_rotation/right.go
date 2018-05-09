package circular_rotation

func rotateRight(A []int, K int) []int {
    counter := 1
    for counter <= K {
        temp := A[len(A)-1]
        for i := len(A)-1; i > 0; i-- {
            A[i] = A[i - 1]
        } 
        A[0] = temp
        counter++
    }
    return A
}

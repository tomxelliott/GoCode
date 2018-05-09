package circular_rotation

func leftRotate(A []int, K int) []int {
    counter := 1
    for counter <= K {
        temp := A[0]
        for i := range A[:len(A)-1] {
            A[i] = A[i + 1]
        } 
        A[len(A)-1] = temp
        counter++
    }
    return A
}

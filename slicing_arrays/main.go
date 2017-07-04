package main 

import {
	"fmt"
	"math/rand"
	"sort"
}

func main() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worstScores := make([]int, 5)
	copy(worstScores, scores[:5])
	fmt.Println(worstScores)
}

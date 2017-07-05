package main 

import(
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("untitled.txt") // fake file (replace with real file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// read the file
	initializedIf()
}

// Cool little tidbit: Initialized If:
func initializedIf() int {
	count := 12
	if x := 10; count > x {
		fmt.Println("Cool!")
	}
	return count
}



// Go may have a Garbage Collector but some resources require that we explicitly release them.
// e.g Close() a file after we are finished with them.
// whatever it is we "defer" will be executed after the enclosing function (main() above)
// helps avoid forgetting about certain things declared within function. 

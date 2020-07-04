package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	// write your code in Go 1.4
	result := 0
	var i int
	for i = 1; i*i < N; i++ {
		if N%i == 0 {
			result += 2
		}
	}
	if i*i == N {
		result++
	}
	return result
}

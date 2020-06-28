package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, A []int) int {
	// write your code in Go 1.4
	var result int
	table := make([]bool, X+1)
	N := (X + 1) * X / 2
	for i := 0; i <= len(A)-1; i++ {
		if !table[A[i]] {
			table[A[i]] = true
			N -= A[i]
		}
		if N == 0 {
			result = i
			break
		}
	}
	if N == 0 {
		return result
	}
	return -1
}

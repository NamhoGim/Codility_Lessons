package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	sort.Ints(A)
	result := A[len(A)-1] * A[len(A)-2] * A[len(A)-3]
	if result < A[0]*A[1]*A[len(A)-1] {
		result = A[0] * A[1] * A[len(A)-1]
	}
	return result
}

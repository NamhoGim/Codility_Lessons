package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	result := 0
	sort.Ints(A)
	for i := range A {
		if i == 0 {
			result++
		} else {
			if A[i] != A[i-1] {
				result++
			}
		}
	}
	return result
}

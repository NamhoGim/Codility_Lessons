package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	N := len(A)
	begin := make([]int, N)
	end := make([]int, N)

	for i := range A {
		begin[i] = i - A[i]
		end[i] = i + A[i]
	}

	sort.Ints(begin)
	sort.Ints(end)

	result := 0
	for i, j := 0, 0; i < N; i++ {
		for j < N && end[i] >= begin[j] {
			result += j - i
			j++
		}
	}
	if result > 10000000 {
		return -1
	}
	return result
}

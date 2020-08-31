package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int) int {
	// write your code in Go 1.4
	result := 0
	s := []int{}
	for i := 0; i < len(A); i++ {
		if B[i] == 0 {
			for len(s) > 0 && s[len(s)-1] < A[i] {
				s = s[:len(s)-1]
			}
			if len(s) == 0 {
				result++
			}
		} else {
			s = append(s, A[i])
		}
	}
	result += len(s)
	return result
}

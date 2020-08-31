package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Solution(A []int) int {
	// write your code in Go 1.4
	check := map[int]bool{}
	N := 0
	for _, n := range A {
		check[n] = true
		N = max(N, n)
	}

	var i int
	for i = 1; i <= N; i++ {
		if _, ok := check[i]; !ok {
			break
		}
	}
	return i
}

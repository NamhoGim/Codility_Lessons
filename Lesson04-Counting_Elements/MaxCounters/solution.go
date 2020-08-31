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

func Solution(N int, A []int) []int {
	// write your code in Go 1.4
	maxCnt := 0
	setAs := 0
	table := make([]int, N)
	for _, q := range A {
		if q >= 1 && q <= N {
			if table[q-1] < setAs {
				table[q-1] = setAs + 1
			} else {
				table[q-1]++
			}
			maxCnt = max(maxCnt, table[q-1])
		} else if q == N+1 {
			setAs = maxCnt
		}
	}
	for i, v := range table {
		if v < setAs {
			table[i] = setAs
		}
	}
	return table
}

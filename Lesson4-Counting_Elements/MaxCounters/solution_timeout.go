package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func maxCounter(table []int, val int) {
	for i := range table {
		table[i] = val
	}
}

func SolutionTimeout(N int, A []int) []int {
	// write your code in Go 1.4
	maxCnt := 0
	table := make([]int, N)
	for _, q := range A {
		if q >= 1 && q <= N {
			table[q-1]++
			maxCnt = max(maxCnt, table[q-1])
		} else if q == N+1 {
			maxCounter(table, maxCnt)
		}
	}
	return table
}

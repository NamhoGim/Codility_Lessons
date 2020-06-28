package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	table := map[int]int{}
	for _, n := range A {
		table[n]++
	}

	var result int
	for k, v := range table {
		if v%2 != 0 {
			result = k
		}
	}
	return result
}

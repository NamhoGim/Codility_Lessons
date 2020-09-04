package solution

import (
	"fmt"
	"testing"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Solution(A []int) int {
	a, b := 1, 0
	N := len(A)
	fibos := []int{}
	for a+b <= N+1 {
		fibos = append(fibos, a+b)
		b, a = a, a+b
	}
	table := make([]int, N+1)
	for i := range table {
		table[i] = 100000
	}

	A = append(A, 1)
	for i, v := range A {
		if v == 1 {
			for _, f := range fibos {
				if i-f == -1 {
					table[i] = 1
					break
				} else if i-f > -1 {
					table[i] = min(table[i], table[i-f]+1)
				}
			}
		}
	}
	if table[N] >= 100000 {
		return -1
	}
	return table[N]
}

func TestSolution(t *testing.T) {
	fmt.Println(Solution([]int{0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0}))
}

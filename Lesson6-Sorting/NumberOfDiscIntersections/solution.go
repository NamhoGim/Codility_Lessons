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

func SolutionOptimal(A []int) int {
	N := len(A)
	intersections := 0
	begin := make([]int, N)
	end := make([]int, N)
	for i := range A {
		begin[i] = i - A[i]
		end[i] = i + A[i]
	}
	sort.Ints(begin)
	sort.Ints(end)

	// Check whether current end point is in between other circle.
	// To make checking faster, check begin point which is located left side of the current end point.
	for i, e := range end {
		n := sort.Search(len(begin), func(i int) bool { return begin[i] > e })
		intersections += n - 1 - i
		if intersections > 10000000 {
			return -1
		}
	}
	return intersections
}

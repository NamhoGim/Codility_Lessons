package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	size, value := 0, 0
	for _, v := range A {
		if size == 0 {
			size++
			value = v
		} else {
			if value != v {
				size--
			} else {
				size++
			}
		}
	}

	var dominiCandidate int
	if size > 0 {
		dominiCandidate = value
	}

	count, dominiIdx := 0, 0
	for i, v := range A {
		if v == dominiCandidate {
			count++
			dominiIdx = i
		}
	}

	if count > len(A)/2 {
		return dominiIdx
	}
	return -1
}

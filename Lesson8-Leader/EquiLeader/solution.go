package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func freqKeyInMap(input map[int]int) int {
	count := 0
	freqKey := 0
	for k, v := range input {
		if count < v {
			count = v
			freqKey = k
		}
	}
	return freqKey
}

func Solution(A []int) int {
	// write your code in Go 1.4
	arrMap := map[int]int{}
	for _, n := range A {
		arrMap[n]++
	}

	k := freqKeyInMap(arrMap)
	if arrMap[k] <= len(A)/2 {
		return 0
	}

	result := 0
	leftCount := 0
	for i, n := range A {
		if n == k {
			leftCount++
		}
		if (leftCount > (i+1)/2) && (arrMap[k]-leftCount > (len(A)-i-1)/2) {
			result++
		}
	}
	return result
}

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

func maxArr(arr []int) int {
	result := 0
	for i, n := range arr {
		if i == 0 {
			result = n
		} else {
			if result < n {
				result = n
			}
		}
	}
	return result
}

func SolutionOptimum(A []int) []int {
	maxA := maxArr(A)
	result := make([]int, len(A))

	elemTable := map[int]int{}
	for _, n := range A {
		elemTable[n]++
	}

	elemDivisors := map[int]map[int]bool{}
	for _, n := range A {
		elemDivisors[n] = map[int]bool{1: true, n: true}
	}

	for d := 2; d*d <= maxA; d++ {
		for elem := d; elem <= maxA; elem += d {
			if _, ok := elemDivisors[elem]; ok {
				if _, ok := elemDivisors[elem][d]; !ok {
					elemDivisors[elem][d] = true
					elemDivisors[elem][elem/d] = true
				}
			}
		}
	}

	for i, n := range A {
		divCount := 0
		for k := range elemDivisors[n] {
			divCount += elemTable[k]
		}
		result[i] = len(A) - divCount
	}

	return result
}

func SolutionTimeOut(A []int) []int {
	// write your code in Go 1.4
	result := make([]int, len(A))
	for i := range result {
		result[i] = len(A) - 1
	}

	var maxVal int
	valueMap := map[int][]int{}
	for i, n := range A {
		if i == 0 {
			maxVal = n
		} else {
			maxVal = max(maxVal, n)
		}
		if _, ok := valueMap[n]; !ok {
			valueMap[n] = []int{i}
		} else {
			valueMap[n] = append(valueMap[n], i)
		}
	}

	for i, n := range A {
		if n != 1 {
			for j := 1; n*j <= maxVal; j++ {
				for _, k := range valueMap[n*j] {
					if k != i {
						result[k]--
					}
				}
			}
		} else {
			for l := range result {
				if l != i {
					result[l]--
				}
			}
		}
	}

	return result
}

package solution

import "testing"

func gcd(x, y int) int {
	for x%y != 0 {
		x, y = y, x%y
	}
	return y
}

func hasSamePrimeDivisors(a, b int) bool {
	gcdAB := gcd(a, b)
	for a > 1 {
		gcdA := gcd(a, gcdAB)
		if gcdA == 1 {
			break
		}
		a /= gcdA
	}
	if a != 1 {
		return false
	}

	for b > 1 {
		gcdB := gcd(b, gcdAB)
		if gcdB == 1 {
			break
		}
		b /= gcdB
	}
	return b == 1
}

func Solution(A []int, B []int) int {
	result := 0
	for i := 0; i < len(A); i++ {
		if hasSamePrimeDivisors(A[i], B[i]) {
			result++
		}
	}
	return result
}

func TestHasSamePrimeDivisors(t *testing.T) {
	if !hasSamePrimeDivisors(15, 75) {
		t.Error("Wrong answer")
	}
}

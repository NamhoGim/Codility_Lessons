package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string, P []int, Q []int) []int {
    // write your code in Go 1.4
    result := make([]int, len(P))
    table := make(map[rune][]int)
    for _, r := range "ACGT" {
        table[r] = make([]int, len(S))
    }
    
    for i, r := range S {
        if i == 0 {
            table[r][i] = 1
        } else {
            table[r][i] = table[r][i-1] + 1
            for k := range table {
                if r != k {
                    table[k][i] = table[k][i-1]
                }
            }
        }
    }
    
    var tmpResult int
    for i := 0; i < len(P); i++ {
        p, q := P[i], Q[i]
        if p == 0 {
            for i, r := range "ACGT" {
                if table[r][q] != 0 {
                    tmpResult = i+1
                    break
                }
            }
        } else {
            for i, r := range "ACGT" {
                if table[r][q] - table[r][p-1] != 0 {
                    tmpResult = i+1
                    break
                }
            }
        }
        result[i] = tmpResult
    }
    return result
}

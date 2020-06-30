package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    minAvgIdx := 0
    minAvgVal := float32((A[0] + A[1])) / 2
    for i := 0; i < len(A)-2; i++ {
        currVal := float32((A[i] + A[i+1])) / 2
        if minAvgVal > currVal {
            minAvgVal = currVal
            minAvgIdx = i
        }
        
        currVal = float32((A[i] + A[i+1] + A[i+2])) / 3
        if minAvgVal > currVal {
            minAvgVal = currVal
            minAvgIdx = i
        }
    }
    if minAvgVal > float32((A[len(A)-1] + A[len(A)-2])) / 2 {
        minAvgIdx = len(A)-2
    } 
    return minAvgIdx
}
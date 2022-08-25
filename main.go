package main

import (
	"fmt"
	"math"
)

func main() {
	n := 4
	fmt.Println(isPowerOfTwo(n))
}
func isPowerOfTwo(n int) bool {
	for i := 0; i < 31; i++ {
		y := float64(i)
		if float64(n) == math.Pow(2, y) {
			return true
		}
	}
	return false
}

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var switches float64 = 100
	root := int(math.Sqrt(switches))
	for i := 1; i < root+1; i++ {
		fmt.Println("switches " + strconv.Itoa(i*i) + " will be on")
	}

	fmt.Println("Total " + strconv.Itoa(root) + " switches will be on in the end out of " + strconv.Itoa(100) + " switches")
}

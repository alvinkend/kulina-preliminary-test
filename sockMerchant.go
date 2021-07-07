package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number of socks")
	fmt.Scan(&n)

	socks := make([]int, 101)

	fmt.Println("Enter the color of each sock")
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)

		socks[temp]++
	}

	count := 0

	for _, value := range socks {
		count = count + value/2
		// proses pengulangan untuk mendapatkan kaos kaki yang sepasang dalam suatu index,
	}

	fmt.Println("output = ", +count)
}

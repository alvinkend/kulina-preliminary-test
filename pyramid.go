package main

import (
	"fmt"
	"strings"
)

func main() {

	var n string
	fmt.Println("Input Number: ")
	fmt.Scan(&n)

	trimString := strings.Replace(n, ".", "", -1)

	result := strings.Split(trimString, "")
	for i := range result {
		fmt.Println(result[i] + strings.Repeat("0", len(result)-i-1))
		//strings.repeat untuk menambahkan angka 0 berdasarkan posisi index
	}
}

package main

import "fmt"

func main() {
	countValley(8, []string{"U", "D", "D", "D", "U", "D", "U", "U"})
}

func countValley(n int, steps []string) {

	if len(steps) != n {
		fmt.Println("Input Salah")
	} else {
		var valley = 0     //nilai awal valley
		var currentPos = 0 //posisi awal daratan.

		for i := 0; i < n; i++ {
			if steps[i] == "U" {
				currentPos++ //posisi bertambah
			} else {
				currentPos-- //posisi berkurang
			}
			if currentPos == 0 && steps[i] == "U" { //ketika posisi di darat, dan dia habis dari dasar laut, maka total valley +1
				valley++
			}
		}
		fmt.Println(valley)
	}
}

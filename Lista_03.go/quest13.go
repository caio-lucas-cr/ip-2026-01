package main

import f "fmt"

func main() {

	var H float64

	f.Println("O valor de H é dado por: ")

	H = 0

	for i := 1; i <= 50; i++ {
		x := 1 + (i-1)*2
		y := 1 + (i-1)*1

		H += float64(x) / float64(y)
	}

	f.Printf("\nH = %.6f", H)
}

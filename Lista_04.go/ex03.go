package main

import f "fmt"

func main() {

	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := len(array)

	f.Printf("Original: %v\n", array)

	inversa(array, 0, n-1)

	f.Printf("Inversa: %v", array)
}

func inversa(array []int, esq int, dir int) {
	if esq >= dir {
		return
	}

	array[esq], array[dir] = array[dir], array[esq]

	inversa(array, esq+1, dir-1)
}

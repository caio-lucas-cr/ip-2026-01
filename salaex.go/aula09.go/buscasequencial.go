package main

import f "fmt"

func main() {
	var lista = []int{1, 45, 7, 4, 98, 47, 17, 14}

	B := BuscaSequencial(lista, 4)

	f.Println(B)

}

func BuscaSequencial(l []int, x int) int {
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			return i
		}
	}
	return -1

}

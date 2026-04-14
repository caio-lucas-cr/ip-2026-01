package main

import f "fmt"

func main() {

	var lista = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	valor := 14
	indice := buscaBinaria(lista, valor)

	if indice != -1 {
		f.Printf("O indíce do valor %d é: %d", valor, indice)
	} else {
		f.Println("O valor não está presente na lista!")
	}

}

func buscaBinaria(l []int, x int) int {
	e := 0
	d := len(l) - 1

	for e <= d {
		m := (e + d) / 2

		if l[m] == x {
			return m
		}

		if l[m] > x {
			d = m - 1
		}

		if l[m] < x {
			e = m + 1
		}
	}
	return -1
}

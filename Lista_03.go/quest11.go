package main

import f "fmt"

func main() {

	var n int

	f.Println("Insira um núemro inteiro positivo: ")
	f.Scan(&n)

	if n <= 0 {
		return
	}

	F := fatorial(n)

	f.Printf("O fatorial de %d é igual a: %d", n, F)
}

func fatorial(n int) int {
	var k int = 1

	for i := 2; i <= n; i++ {
		k = k * i
	}
	return k

}

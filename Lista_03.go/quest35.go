package main

import (
	f "fmt"
	"strconv"
)

func main() {
	var n int
	var base2 string

	f.Println("Insira um número na base 10 para transformá-lo na base 2: ")
	f.Scan(&n)

	orign := n

	if n < 0 {
		f.Println("Número inválido!")
		return
	}

	if n == 0 {
		base2 = "0"
	}

	for n > 0 {
		resto := n % 2
		base2 = strconv.Itoa(resto) + base2
		n = n / 2
	}

	f.Printf("O número %d na base 2 é: %s", orign, base2)
}

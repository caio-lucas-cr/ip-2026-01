package main

import (
	f "fmt"
	"strconv"
)

func main() {
	var n int
	var base16, digito string

	f.Println("Insira um número na base 10 para transformá-lo na base 16: ")
	f.Scan(&n)

	orign := n

	if n < 0 {
		f.Println("Número inválido!")
		return
	}

	if n == 0 {
		base16 = "0"
	}

	for n > 0 {
		resto := n % 16
		if resto >= 10 {
			digito = string(rune(resto + 55))
		} else {
			digito = strconv.Itoa(resto)
		}

		base16 = digito + base16
		n = n / 16
	}

	f.Printf("O número %d na base 16 é: %s", orign, base16)
}

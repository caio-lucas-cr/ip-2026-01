package main

import (
	f "fmt"
	sc "strconv"
)

func main() {

    var numero int
    
	f.Println("Insira um número na forma decimal para transformá-lo em binário: ")
    f.Scan(&numero)
	
	if numero == 0 {
		f.Printf("O binário de %d é: 0\n", numero)
	} else {
		resultado := DecimalParaBinario(numero)
		f.Printf("O binário de %d é: %s\n", numero, resultado)
	}
}

func DecimalParaBinario(n int) string {
	if n == 0 {
		return ""
	}

	return DecimalParaBinario(n/2) + sc.Itoa(n%2)
}
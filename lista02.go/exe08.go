package main

import f "fmt"

func main() {

	var n int

	f.Println("Digite um número :")
	f.Scan(&n)

	if (n>20) && (n<90) {
		f.Println("O número ESTÁ compreendido entre 20 e 90")
	} else {
		f.Println("O número NÃO ESTÁ compreendido entre 20 e 90")
	}
}
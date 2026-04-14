package main

import f "fmt"

func main() {

	var s, m, qtd int

	f.Println("\nO valor da soma e média dos números pares entre 50 e 70 é dado por: \n")

	for i := 50; i <= 70; i++ {
		if i%2 == 0 {
			s += i
			qtd++
		}
	}

	m = s / qtd

	f.Println("A soma é: ", s)
	f.Println("A média aritmética é: ", m)

}

package main

import f "fmt"

func main() {

	var s int

	f.Println("Imprindo os números de 1 a 20, e a soma deles: \n")

	for i := 1; i <= 20; i++ {
		f.Println(i)
		s += i
	}

	f.Println("Soma = ", s)

}

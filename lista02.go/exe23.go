package main

import f "fmt"

func main() {

	var idade int
	var class_eleitoral string

	f.Println("Digite sua idade: ")
	f.Scan(&idade)

	if idade < 16 {
		class_eleitoral = "Não-eleitor"
	} else if idade >= 18 && idade <= 65 {
		class_eleitoral = "Eleitor obrigatório"
	} else if (idade >= 16 && idade < 18) || idade > 65 {
		class_eleitoral = "Eleitor facultativo"
	}
	f.Println("Classe Eleitoral: ", class_eleitoral)

}
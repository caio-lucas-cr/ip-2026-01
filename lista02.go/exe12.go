package main

import f "fmt"

func main() {

	var id int
	var clas string

	f.Println("Insira a idade da pessoa: ")
	f.Scan(&id)

	if id >= 0 && id <= 2 {
		clas = "Recém-nascido"
		f.Println("Classificação da pessoa: ", clas)
	} else if id > 2 && id <= 11 {
		clas = "Criança"
		f.Println("Classificação da pessoa: ", clas)
	} else if id > 12 && id <= 19 {
		clas = "Adolescente"
		f.Println("Classificação da pessoa: ", clas)
	} else if id > 19 && id <= 55 {
		clas = "Adulto"
		f.Println("Classificação da pessoa: ", clas)
	} else if id > 55 {
		clas = "Idoso"
		f.Println("Classificação da pessoa: ", clas)
	} else {
		f.Println("Idade inválida!")
	}
}

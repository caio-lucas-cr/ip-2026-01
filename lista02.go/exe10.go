package main

import f "fmt"

func main() {

	var dest, iv string
	var val float32

	f.Println("Digite o destino da viagem: ")
	f.Scan(&dest)

	f.Println("DIgite se a viagem é de ida e volta, ou de só ida: ")
	f.Scan(&iv)

	if dest == "1" {
		if iv == "1" {
			val = 900
			f.Printf("O preço da passagem é de R$ %.2f", val)
		} else if iv == "2" {
			val = 500
			f.Printf("O preço da passagem é de R$ %.2f", val)
		}
	} else if dest == "2" {
		if iv == "1" {
			val = 650
			f.Printf("O preço da passagem é de R$ %.2f", val)
		} else if iv == "2" {
			val = 350
			f.Printf("O preço da passagem é de R$ %.2f", val)
		}
	} else if dest == "3" {
		if iv == "1" {
			val = 600
			f.Printf("O preço da passagem é de R$ %.2f", val)
		} else if iv == "2" {
			val = 350
			f.Printf("O preço da passagem é de R$ %.2f", val)
		}

	} else if dest == "4" {
		if iv == "1" {
			val = 550
			f.Printf("O preço da passagem é de R$ %.2f", val)
		} else if iv == "2" {
			val = 300
			f.Printf("O preço da passagem é de R$ %.2f", val)
		}

	} else {
		f.Println("Destino ou tipo de viagem inválido")
	}
}

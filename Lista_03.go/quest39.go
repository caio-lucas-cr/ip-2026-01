package main

import f "fmt"

func main() {

	const totalBois = 90

	var id, idGordo, idMagro int
	var peso, pesoGordo, pesoMagro float64

	f.Println("--- Sistema de Controle de Peso do Frigorífico ---")

	for i := 1; i <= totalBois; i++ {
		f.Printf("\nEntrada de dados do boi %d:\n", i)
		f.Println("Número de identificação: ")
		f.Scan(&id)
		f.Println("Peso (kg): ")
		f.Scan(&peso)

		if i == 1 {
			pesoGordo = peso
			idGordo = id
			pesoMagro = peso
			idMagro = id
			continue
		}

		if peso > pesoGordo {
			pesoGordo = peso
			idGordo = id
		}

		if peso < pesoMagro {
			pesoMagro = peso
			idMagro = id
		}
	}

	f.Println("\n" + "--- RESULTADO FINAL ---")
	f.Printf("Boi mais GORDO: ID %d com %.2f kg\n", idGordo, pesoGordo)
	f.Printf("Boi mais MAGRO: ID %d com %.2f kg\n", idMagro, pesoMagro)
}

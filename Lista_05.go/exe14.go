package main

import f "fmt"

func main() {
	const tam = 10
	var v1 [tam]int
	var v2 [tam]int
	var resultante [tam * 2]int

	f.Println("Digite os 10 elementos do Vetor 1:")
	for i := 0; i < tam; i++ {
		f.Scan(&v1[i])
	}

	f.Println("Digite os 10 elementos do Vetor 2:")
	for i := 0; i < tam; i++ {
		f.Scan(&v2[i])
	}

	j := 0
	for i := 0; i < tam; i++ {
		resultante[j] = v1[i]   
		resultante[j+1] = v2[i] 
		j += 2                  
	}

	f.Println("\nVetor 1:", v1)
	f.Println("Vetor 2:", v2)
	f.Println("Vetor Intercalado:", resultante)
}
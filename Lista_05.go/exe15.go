package main

import f "fmt"

func main() {

	vetor:=make([]int, 30)
	vetorFinal:=make([]int, 30)

	for i:=0; i<len(vetor); i++{
		f.Printf("Vetor[%d]: ",i)
		f.Scan(&vetor[i])
	}

	for i:=0; i<len(vetor); i++{
		if i%2==0 {
			vetorFinal[i]=2*vetor[i]
		} else {
			vetorFinal[i]=3*vetor[i]
		}
	}

	f.Println("Vetor: ", vetor)
	f.Println("Vetor final: ", vetorFinal)
}
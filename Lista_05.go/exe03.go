package main

import f "fmt"

func main() {
    var somaPares, qtdImpares int
    vetor := make([]int, 10)

    f.Println("Insira 10 números inteiros:")

    for i := 0; i < len(vetor); i++ {
        f.Printf("Vetor [%d]: ", i)
        f.Scan(&vetor[i])
    }

    var pares []int
    var impares []int

    for i := 0; i < len(vetor); i++ {
        numero := vetor[i]
        if numero%2 == 0 {
            pares = append(pares, numero)
            somaPares += numero
        } else {
            impares = append(impares, numero)
            qtdImpares++
        }
    }

    f.Println("\n--- RESULTADOS ---")
    
    f.Printf("Números pares: %v\n", pares)
    f.Printf("A soma dos números pares é: %d\n", somaPares) 
    
    f.Printf("\nNúmeros ímpares: %v\n", impares)
    f.Printf("A quantidade de números ímpares é: %d\n", qtdImpares)
}
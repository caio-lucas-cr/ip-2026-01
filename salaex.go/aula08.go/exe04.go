package main

import f "fmt"

func main() {

	 var n int

	 f.Println("Insira um núemro inteiro: ")
	 f.Scan(&n)

	 F:=fatorial(n)

	 f.Printf("O fatorial de %d é igual a: %d", n, F)
}

func fatorial(n int)int {
	var k int = 1

	for i:=1 ; i <=n; i++ {
		k=k*i
	} 
	return k

}
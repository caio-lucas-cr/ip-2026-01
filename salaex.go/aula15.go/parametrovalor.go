package main
import "fmt"
func main() {
     v := 0
     vf := 0
     vf = incrementa(v)
     fmt.Println(v) // 0
     fmt.Println(vf) // 1
}
func incrementa(x int) int {
     return x + 1 
}


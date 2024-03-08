package main

import (
  "fmt"
)

func hanoi(n int, origem, destino, auxiliar string) {
  if n == 1 {
    fmt.Printf("Mover disco 1 do %s para %s   (n===1) \n\n", origem, destino)
    return
  }
  
  hanoi(n-1, origem, auxiliar, destino)
  
  fmt.Printf("Mover disco %d do %s para %s   (n!!!1) \n", n, origem, destino)
  fmt.Printf("primeira chamada\n\n")
  hanoi(n-1, auxiliar, destino, origem)
  fmt.Printf("segunda chamada\n\n")
}

func main() {
  var discos int
  fmt.Println("Quantos discos tem a torre?")
  fmt.Scanf("%d", &discos)
  hanoi(discos, "A", "C", "B")
}

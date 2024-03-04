package main

import (
  "fmt"
  "math/rand"
  "math"
  "./geometria"
)

type Ponto struct {
    X, Y float64
}


func main() {
  
  imprimeVetor(10)

  
  fmt.Println("---Ex:2---")
  var frase string
  fmt.Println("Digite uma string")
  fmt.Scanln(&frase)
  fmt.Println("String invertida:", inverterString(frase))

  
  fmt.Println("---Ex:3---")
  ponto := Ponto{5, 8}
  fmt.Println("Distância até a origem:", ponto.DistanciaOrigem())

  
  fmt.Println("---Ex:4---")
  var base, altura float64
  fmt.Println("Insira a base do retângulo:")
  fmt.Scanln(&base)
  fmt.Println("Insira a altura do retângulo:")
  fmt.Scanln(&altura)

  fmt.Printf("Área do retângulo: %.2f\n", geometria.AreaRetangulo(base, altura))
  fmt.Printf("Perímetro do retângulo: %.2f\n", geometria.PerimetroRetangulo(base, altura))

  
  
}


func imprimeVetor(tamanho int) {
  vetor := make([]int, tamanho)

  for i := 0; i < tamanho; i++ {
      vetor[i] = rand.Intn(100)
      fmt.Println(vetor[i])
  }

}


func inverterString(input string) string {
  caracteres := []rune(input)

  for i, j := 0, len(caracteres)-1; i < j; i, j = i+1, j-1 {
    fmt.Println(i, j)
    caracteres[i], caracteres[j] = caracteres[j], caracteres[i]
  }

  return string(caracteres)
}


func (p Ponto) DistanciaOrigem() float64 {
  return math.Sqrt((p.X*p.X) + (p.Y*p.Y))
}


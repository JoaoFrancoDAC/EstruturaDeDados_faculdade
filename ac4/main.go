package main

import "fmt"

func parArray(array []int, desejado int) (int, int) {
  ponteiro_esquerdo := 0
  ponteiro_direito := len(array) - 1

  for ponteiro_esquerdo < ponteiro_direito {
    soma := array[ponteiro_esquerdo] + array[ponteiro_direito]

    if soma == desejado {
      return ponteiro_esquerdo, ponteiro_direito
    } else if soma < desejado {
      ponteiro_esquerdo++
    } else {
      ponteiro_direito--
    }
  }

  return -1, -1
}

func main() {
  array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  alvo := 21
  menorPonteiro, maiorPonteiro := parArray(array, alvo)

  if menorPonteiro == -1 && maiorPonteiro == -1 {
    fmt.Println("Nenhum par encontrado")
  } else {
    menorValor := array[menorPonteiro]
    maiorValor := array[maiorPonteiro]
    fmt.Printf("Par de números encontrado primeiro: (%d, %d)\n", menorValor, maiorValor)
  }
}

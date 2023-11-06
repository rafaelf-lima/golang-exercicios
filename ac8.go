package main

import (
	"fmt"
)

type No struct {
  Chave    int
  Esq, Dir *No
  Altura   int
}

type Arvore struct {
  Raiz *No
}

func main() {
  a := &Arvore{}
  
  resultado := busca(5, a.Raiz)
  fmt.Println(resultado)

  a.Raiz = insere(5, a.Raiz)
  resultado2 := busca(5, a.Raiz)
  fmt.Println(resultado2)

  resultado3 := busca(2, a.Raiz)
  fmt.Println(resultado3)

  a.Raiz = insere(10, a.Raiz)
  resultado4 := busca(10, a.Raiz)
  fmt.Println(resultado4)

  resultado5 := busca(44, a.Raiz)
  fmt.Println(resultado5)

  a.Raiz = insere(10, a.Raiz)
  
}

  

func busca(valor int, no *No) int {
  if no == nil { return 0 }
  if no.Chave == valor { return 1 }
  if valor < no.Chave {
    if no.Esq == nil { return 2 }
    no = no.Esq
    return busca(valor, no)
  } else {
    if no.Dir == nil { return 3 }
    // no = no.Dir
    return busca(valor, no.Dir)
  }
}

func insere(valor int, no *No) *No {
  if no != nil {
    if valor > no.Chave {
      no.Dir = insere(valor, no.Dir)
    } else if valor < no.Chave {
      no.Esq = insere(valor, no.Esq)
    } else {
      fmt.Println("Número já inserido!")
    }
    return no
  }
  return &No{Chave: valor, Esq: nil, Dir: nil}
}


package main

import (
	"fmt"
)

type No struct {
  Valor int
  Prox *No
}

type Lista struct {
  Cab *No
}


func main() {
  lista := &Lista{}
  lista.insereComeco(1)
  lista.insereComeco(2)
  lista.insereComeco(3)
  lista.insereComeco(4)
  lista.insereComeco(7)
  fmt.Println("---------------")
  lista.exibeNos()
  lista.removeComeco()
  fmt.Println("---------------")
  lista.exibeNos()
  lista.removeComeco()
  fmt.Println("---------------")
  lista.exibeNos()
  lista.removeComeco()
  fmt.Println("---------------")
  lista.exibeNos()
  
}

func (l *Lista) insereComeco(valor int) {
  novoNo := &No{Valor: valor, Prox: nil}
  if l.Cab == nil {
    novoNo.Prox = novoNo
    l.Cab = novoNo
  } else {
    novoNo.Prox = l.Cab.Prox
    l.Cab.Prox = novoNo
    l.Cab = novoNo
  }
}

func (l *Lista) exibeNos(){
  no := l.Cab
  if no == nil {
    fmt.Println("Lista vazia.")
  }  else {
    for {
      fmt.Println(no.Valor)
      no = no.Prox
      if no == l.Cab {
        return
      }
    }
  }
}


func (l *Lista) removeComeco() {
  no := l.Cab
  if no == nil {
    fmt.Println("Lista vazia.")
    return
  }
  if no.Prox == l.Cab {
    l.Cab = nil
  } else {
    for no.Prox != l.Cab {
      no = no.Prox
    }

    l.Cab = l.Cab.Prox
    no.Prox = l.Cab
  }
}

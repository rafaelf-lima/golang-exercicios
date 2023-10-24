package main

import (
	"fmt"
)
type No struct {
  Valor int
  Prox *No
  Ant *No
}

type Lista struct {
  Cab *No
}


func main() {
  lista := Lista{}
  lista.insereOrd(4)
  lista.insereOrd(1)
  lista.insereOrd(2)
  lista.insereOrd(5)
  lista.exibe()
  fmt.Println(lista.busca(2))
}

func (l *Lista) insereOrd(valor int) {
    novoNo := &No{Valor: valor, Prox: nil, Ant: nil}

    if l.Cab == nil {
        l.Cab = novoNo
        novoNo.Prox = l.Cab
        novoNo.Ant = l.Cab
    } else if valor < l.Cab.Valor {
        novoNo.Prox = l.Cab
        novoNo.Ant = l.Cab.Ant
        l.Cab.Ant.Prox = novoNo
        l.Cab.Ant = novoNo
        l.Cab = novoNo
    } else {
        no := l.Cab.Prox
        for no != l.Cab && valor >= no.Valor {
            no = no.Prox
        }
        novoNo.Prox = no
        novoNo.Ant = no.Ant
        no.Ant.Prox = novoNo
        no.Ant = novoNo
    }
}

func (l *Lista) exibe() {
no := l.Cab
for no != nil {
    fmt.Println(no.Valor)
    no = no.Prox
    if no == l.Cab {
        return
    }
}
}

func (l *Lista) busca(valor int) *No {
  no := l.Cab
  for no != nil {
    if no.Valor == valor {
      return no
    }
    no = no.Prox
  }
  return nil
}

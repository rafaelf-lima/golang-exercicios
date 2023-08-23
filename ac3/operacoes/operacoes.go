package operacoes

import (
	"ac3/contato"
	"fmt"
)

func AdicionaContato(ctt contato.Contato, a *[5]contato.Contato){
	for ind, c := range a {
		if (c == contato.Contato{}) {
			a[ind] = ctt
			break
		}
	}
}


func ExcluiContato(a *[5]contato.Contato){
	if (a[0] == contato.Contato{}) {
		fmt.Println("Lista de contatos vazia!")
	}

	for ind, c := range a {
		if (c == contato.Contato{}) {
			a[ind-1] = contato.Contato{}
		}
	}
}
#### Lista duplamente encadeadas

Uma lista duplamente encadeada é um modelo de estrutura de coleção de dados que cada nó possui um ponteiro para o nó anterior, assim como também aponta para nó sucessor. Ele difere da lista unicamente encadeada justamente por também apontar ao nó anterior, permitindo percorrer a lista em ambas as direções.

Para identificar o início e final da lista, há de perceber que se nó anterior é nulo, o nó sucessor a esse é o primeiro da lista, uma vez que nada está antes dele, mas se o nó sucessor for nulo, o nó anterior será o último da lista, já que não haverá nada após esse elemento. A presença de dois ponteiros em cada nó permite operações de inserção e remoção de modo eficiente ao longo da lista, além de dinâmica, possibilitando a mudança no tamanho também. Permite, também, a navegação nos dois sentidos potencializando diversos algoritmos.

A ordenação da lista duplamente encadeada otimiza a estrutura ainda mais, que já possui a capacidade de ter ponteiros apontando para os elementos anteriores e sucessores, permitindo navegar pela lista a partir de diferentes direções. Desse modo, assim como em outras estruturas que estão ordenadas, a lista duplamente encadeada ordenada traz performance para a estrutura e possibilidade de operações com os dados presentes na lista.

Para a inserção do nó no final da lista, o seguinte pseudocódigo é válido:
```
FUNÇÃO INSERE_ORDENADO(lista, valor)
  novoNo := Cria nó
  SE lista.cab == NULL ENTÃO
    lista.cab = novoNo
    novoNo.prox = lista.cab
    novoNo.ant = lista.
  SE valor < lista.cab.valor ENTÃO
    novoNo.prox = lista.cab
    novoNo.ant = lista.cab.ant
    lista.cab.ant.prox = novoNo
    lista.cab.ant = novoNo
    lista.cab = novoNo
  SENÃO
    no = lista.cab.prox
    ENQUANTO no != lista.cab e valor >= no.valor FAÇA
      no = no.prox
    novoNo.prox = no
    novoNo.ant = no.ant
    no.ant.prox = novoNo
    no.ant = novoNo
FIM
``` 


No caso da exibição, um dos pseudocódigos válidos é:
```
PROGRAMA EXIBE(lista)
  no = lista.cab
  ENQUANTO no != NULL FAÇA
    Escreve valor do no
    no = no.prox
  no = lista.cab.ant
  ENQUANTO no != lista.cab FAÇA
    Escreve valor do no
    no = no.ant
FIM
```

Por fim, para a busca, um pseudocódigo para esse problema é:
```
FUNÇÃO BUSCA(lista, valor)
  no := lista.cab
  ENQUANTO no != nulo FAÇA
    SE no.valor == valor ENTÃO
      RETORNA no
    no = no.prox
  RETORNA nulo
FIM
```

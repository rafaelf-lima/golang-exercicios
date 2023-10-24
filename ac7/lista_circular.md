#### Listas circulares simplesmente encadeadas
Uma lista circular simplesmente encadeada é uma estrutura que gera uma coleção de dados. Nesse sentido, pode-se dizer que a lista circular tem semelhanças com a lista encadeada simples, porém com diferenças fundamentais entre elas. A principal diferença é o conceito de que o último elemento na lista circular não aponta para _null_, como acontece lista encadeada; esse _null_ é útil para apontar o ponto final da lista até aquele momento. Desse modo, o último elemento da lista circular aponta para o primeiro elemento, por isso, há o nome de "lista circular", com um nó sempre apontando para outro nó. 

Como o último elemento não aponta para _null_, a lista circular é útil para estruturas que requerem iterações sucessivas, já que o último elemento, na verdade, aponta para o primeiro elemento da lista e, então, permite que percorra continuamente a lista diversas vezes de modo eficiente. Assim, não há a presença de _null_ neste modelo de lista. 

As principais operações para lista circular simplesmente encadeada são a inserção, remoção e exibição dos nós. No caso da inserção, o algoritmo para inserção no início da lista é necessário criar um nó para armazenar os valores; caso a lista esteja vazio, o nó a ser inserido será o primeiro e também o último, porém, caso a lista já esteja preenchida o novo nó será adicionado no início.

Segundo o site tutorialspoint.com, os passos para inserção no início da lista são: 
```
1. INÍCIO
2. Verifique se a lista está vazia
3. Se a lista estiver vazia, adicione o nó e faça o ponteiro da cabeça apontar para esse nó
4. Se a lista não estiver vazia, ligue o nó existente na cabeça como o próximo nó para o novo nó.
5. Faça o novo nó ser a nova cabeça.
6. FIM
```

Em pseudocódigo, um exemplo a seguir pode ser:
```
PROGRAMA INSERE(lista, dado)
  novo_no := cria_no
  SE lista.cab == NULL ENTÃO
    lista.cab := novo_no
    lista.cab.prox := lista.cab
  SENÃO
    novo_no.prox := lista.cab
    lista.cab := novo_no
FIM
  ```


No caso da exibicão dos nós de uma lista circular simplesmente encadeada, basta atribuir uma variável ao último elemento da lista; caso seja vazio, a lista também estará vazio, uma vez que cada nó da lista deve apontar para um outro nó, porém, caso seja diferente de _null_, cria-se um laço de repetição para escrever os elementos da lista até chegar ao último elemento

```
PROGRAMA EXIBE(lista)
    no := lista.cab
    SE lista.cab == NULL ENTÃO
      ESCREVER "A lista está vazia"
      FIM
    no := lista.cab
    ENQUANTO no != NULL FAÇA:
      ESCREVER no.dado
      no := no.prox
 FIM
```

Para remoção do primeiro item da lista circular, deve-se verificar se a lista está vazia, caso contrário basta atribuir o endereço do próximo nó ao início da lista e se estiver apenas um elemento a lista vazia.

```
PROGRAMA REMOVE(lista)
  SE lista.cab == NULL ENTÃO
    RETORNA
  SENÃO
    lista.cab := lista.cab.prox
    SE lista.cab == lista.cab.prox ENTÃO
      lista.cab := NULL
FIM
```

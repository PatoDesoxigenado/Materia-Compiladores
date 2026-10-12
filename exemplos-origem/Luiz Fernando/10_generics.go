package main

import "fmt"

// Pilha genérica que aceita qualquer tipo (T any)
type Pilha[T any] struct {
	itens []T
}

// Push empilha um item do tipo T
func (p *Pilha[T]) Push(item T) {
	p.itens = append(p.itens, item)
}

// Pop desempilha o último item e retorna (ou retorna valor zero e falso se vazia)
func (p *Pilha[T]) Pop() (T, bool) {
	if len(p.itens) == 0 {
		var zero T
		return zero, false
	}
	topo := len(p.itens) - 1
	item := p.itens[topo]
	p.itens = p.itens[:topo]
	return item, true
}

type Usuario struct {
	ID   int
	Nome string
}

func main() {
	fmt.Println("=== EXEMPLO 10: GENERICS ===")

	// Pilha de Inteiros
	pilhaInt := Pilha[int]{}
	pilhaInt.Push(10)
	pilhaInt.Push(20)
	valInt, _ := pilhaInt.Pop()
	fmt.Printf("Pop em Pilha de int: %d\n", valInt) // Saída: 20

	// Pilha de Strings
	pilhaString := Pilha[string]{}
	pilhaString.Push("A")
	pilhaString.Push("B")
	valStr, _ := pilhaString.Pop()
	fmt.Printf("Pop em Pilha de string: %s\n", valStr) // Saída: B

	// Pilha de Structs (Usuario)
	pilhaUsers := Pilha[Usuario]{}
	pilhaUsers.Push(Usuario{ID: 1, Nome: "Alice"})
	pilhaUsers.Push(Usuario{ID: 2, Nome: "Bob"})
	valUser, _ := pilhaUsers.Pop()
	fmt.Printf("Pop em Pilha de Struct: %+v\n", valUser) // Saída: {ID:2 Nome:Bob}
}

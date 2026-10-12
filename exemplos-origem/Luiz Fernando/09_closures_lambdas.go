package main

import (
	"fmt"
)

// geradorDeTickets retorna uma função anônima (closure) que retorna inteiros. Esta função "lembra" do ambiente em que foi criada, ou seja, ela mantém a variável 'contador' viva em memória mesmo após 'geradorDeTickets' ter retornado.
func geradorDeTickets() func() int {
	contador := 0

	return func() int {
		contador++
		return contador
	}
}

// aplicarDesconto é uma função que recebe um comportamento (lambda) como parâmetro.
func aplicarDesconto(preco float64, calculoDesconto func(float64) float64) float64 {
	return calculoDesconto(preco)
}

func main() {
	fmt.Println("=== EXEMPLO 9: CLOSURES E LAMBDAS ===\n")

	fmt.Println("--- Sistema de Senhas (Tickets) ---")

	gerarFilaComum := geradorDeTickets()
	gerarFilaPreferencial := geradorDeTickets()

	fmt.Printf("Ticket Comum 1: %d\n", gerarFilaComum()) // 1
	fmt.Printf("Ticket Comum 2: %d\n", gerarFilaComum()) // 2
	fmt.Printf("Ticket Comum 3: %d\n", gerarFilaComum()) // 3

	fmt.Printf("Ticket Preferencial 1: %d\n", gerarFilaPreferencial()) // 1
	fmt.Printf("Ticket Preferencial 2: %d\n", gerarFilaPreferencial()) // 2

	fmt.Println("\n--- Sistema de Descontos (Comportamento Injetado) ---")
	precoProduto := 100.0

	descontoDezPorcento := func(valor float64) float64 {
		return valor - (valor * 0.10)
	}

	precoFinal := aplicarDesconto(precoProduto, descontoDezPorcento)
	fmt.Printf("Preço original: R$ %.2f\n", precoProduto)
	fmt.Printf("Preço com 10%% de desconto: R$ %.2f\n", precoFinal)

	// Podemos passar uma lambda diretamente na chamada da função
	precoFinalBlackFriday := aplicarDesconto(precoProduto, func(v float64) float64 {
		return v - (v * 0.50) // 50% de desconto
	})

	fmt.Printf("Preço Black Friday (50%% de desconto): R$ %.2f\n", precoFinalBlackFriday)
}

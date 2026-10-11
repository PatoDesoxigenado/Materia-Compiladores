/*
- Não foram usadas IAs neste exemplo.

Exemplo 3: Ponteiros e Passagem de Parâmetros
- Arquivo: 03_ponteiros_referencias.go
- Tópicos: Passagem por valor vs. simulação de referência usando ponteiros explícitos (*T e &x).
- Ideia de Implementação: Criar um sistema de "Transação Bancária", demonstrando que tentar alterar
o saldo passando a struct por valor não funciona, exigindo o uso de ponteiros para efetivar o depósito/saque.
*/

package main

import "fmt"

type ContaBancaria struct {
	Titular string
	Numero  int
	Saldo   float64
}

// Transações por valor: cópiam o valor. Valor original não muda.
func depositoValor(conta ContaBancaria, valor float64) {
	if valor <= 0 {
		fmt.Printf("O depósito deve ser maior que 0.\n")
		return
	}
	conta.Saldo += valor
	fmt.Printf("[Função de deposito por Valor] Saldo: R$ %.2f \n", conta.Saldo)
}

func saqueValor(conta ContaBancaria, valor float64) {
	if valor <= 0 || valor > conta.Saldo {
		fmt.Printf("Não é póssivel realizar esse saque.\n")
		return
	}
	conta.Saldo -= valor
	fmt.Printf("[Função de saque por Valor] Saldo: R$ %.2f \n", conta.Saldo)
}

// Transações por ponteiros: Recebem o endereço da memória. Valor original muda.
func depositoPonteiro(conta *ContaBancaria, valor float64) {
	if valor <= 0 {
		fmt.Printf("O depósito deve ser maior que 0. \n")
		return
	}
	conta.Saldo += valor
	fmt.Printf("[Função de deposito por Ponteiro] Saldo: R$ %.2f \n", conta.Saldo)
}

func saquePonteiro(conta *ContaBancaria, valor float64) {
	if valor <= 0 || valor > conta.Saldo {
		fmt.Printf("Não é póssivel realizar esse saque.\n")
		return
	}
	conta.Saldo -= valor
	fmt.Printf("[Função de saque por Ponteiro] Saldo: R$ %.2f \n", conta.Saldo)
}

func main() {
	conta := ContaBancaria{
		Titular: "Maria",
		Numero:  1234,
		Saldo:   5000.00,
	}

	fmt.Println("=== Sistema Bancário ===")
	fmt.Printf("Titular: %s\n", conta.Titular)
	fmt.Printf("Número da conta: %d\n", conta.Numero)
	fmt.Printf("Saldo inicial: R$ %.2f\n\n", conta.Saldo)

	fmt.Println("=== Teste Depósito ===")

	depositoValor(conta, 100)
	fmt.Printf("Saldo após função por valor: R$ %.2f\n\n", conta.Saldo)

	depositoPonteiro(&conta, 100)
	fmt.Printf("Saldo após por ponteiro: R$ %.2f\n\n", conta.Saldo)

	fmt.Println("=== Teste Saque ===")

	saqueValor(conta, 100)
	fmt.Printf("Saldo após função por valor: R$ %.2f\n\n", conta.Saldo)

	saquePonteiro(&conta, 100)
	fmt.Printf("Saldo após função por Ponteiro: R$ %.2f\n\n", conta.Saldo)

}

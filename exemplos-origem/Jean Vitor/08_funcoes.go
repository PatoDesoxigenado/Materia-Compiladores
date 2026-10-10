/*
  - Não foram usadas IAs neste exemplo.

Exemplo 8: Funções (Retornos e Variádicos)
  - Arquivo: `08_funcoes.go`
  - Tópicos: Múltiplos retornos, retornos nomeados e parâmetros variádicos.
  - Ideia de Implementação: Uma função "Calculadora de Estatísticas" que recebe uma quantidade infinita de notas (variádico), e retorna três valores: o total, a média e um erro (caso nenhuma nota seja enviada).
*/
package main

import (
	"fmt"
	"errors"
	"strconv"
	"os"
)

// função variádica
func total(notas ...float64) float64 {
	resultado := .0

	for _, nota := range notas {
		resultado += nota
	}
	fmt.Println("Somatório total das notas: ", resultado)
	return resultado
}

func média(notas []float64) float64 {
	resultado := total(notas...) / float64(len(notas))
	
	fmt.Println("Média das notas: ", resultado)
	return resultado
}

// função com valor de retorno nomeado
func commandLineToList() (notas []float64, err error) {
	commandLine := os.Args[1:]

	for _, item := range commandLine {
		aux, err := strconv.ParseFloat(item, 64)
		
		if err != nil {
			return notas, err
				
		} else {
			notas = append(notas, aux)
		}
	}
	return notas, err
}

func run() error {
	notas, err := commandLineToList()
	
	if len(notas) == 0 {
		err = errors.New("Parâmetros inválidos ou faltando.")
	
	} else {
		média(notas)
	}
	return err
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %s\n", err)
		os.Exit(1)
	}
}

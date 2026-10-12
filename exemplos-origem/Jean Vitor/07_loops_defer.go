package main

import (
	"fmt"
	"os"
	"bufio"
	"errors"
	"strconv"
)

// presença de for infinito | while (true) {...} |
func buscador() string {
	var local string
	
	for {
		fmt.Print("Insira o local do arquivo .log: ")
		fmt.Scanln(&local)
		
		_, err := os.Open(local)
		if err != nil {
			fmt.Println("Local ou arquivo inválido. Tente novamente.")
		
		} else {
			return local
		}
	}
}

// presença de for condicional | while (scanner.Scan()) {...} | e defer
func scanArquivo() ([]string, error) {
	var linhas []string
	
	arquivo, _ := os.Open(buscador())
	defer arquivo.Close() // será executado só depois do resto
	
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}
    
	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler arquivo:", err)
		return linhas, err
	
	} else {
		return linhas, nil
	}
}

func commandLineN() (int, error) {	
	n, err := strconv.Atoi(os.Args[1])
		
	if err != nil {
		return n, err
			
	} else {
		return n, err
	}
}

// presença de for clássico | for (int i = 15; i < 18; i++) {...} |
func últimasNLinhas(n int) ([]string, error) {
	var resultado []string 
	linhas, err := scanArquivo()
	
	if (len(linhas) - n < 0) || ((len(linhas) - n > len(linhas))) {
		return nil, errors.New("Entrada fora do intervalo aceitável. Quantidade de linhas no arquivo: " + strconv.Itoa(len(linhas)))
	}
	
	for i := len(linhas) - n; i < len(linhas); i++ {
		resultado = append(resultado, linhas[i])
	}
	return resultado, err
} 

// presença de for com range | for i in resultado: ... |
func run() error {
	if len(os.Args) < 2 {
		return errors.New("Nenhuma entrada passada.")
	}
	
	índice, err := commandLineN()
	
	if err != nil {
		return errors.New("Entrada não é um número inteiro.")
	}
	resultado, err := últimasNLinhas(índice)
	
	if err == nil {
		for _, linha := range resultado {
    		fmt.Println("==========")
    		fmt.Println(linha)
    		fmt.Println("==========")
    	}
    	return nil
    
    } else {
		return err
	}
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %s\n", err)
		os.Exit(1)
	}
}

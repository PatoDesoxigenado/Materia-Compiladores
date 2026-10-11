package main

import (
	"fmt"
	"time"
)

// APIFake: simula a chamada de um servidor externo
func APIFake(nome string, tempoDeResposta time.Duration, canal chan<- string) {
	time.Sleep(tempoDeResposta) //simula a latência de rede
	canal <- fmt.Sprintf("Resposta bem-sucedida da %s", nome)
}

func main() {
	// Canais para receber as respostas das APIs
	canalAPI1 := make(chan string)
	canalAPI2 := make(chan string)

	fmt.Println("Iniciando requisições em paralelo")

	// API 1 é muito rápida (500ms), API 2 é lenta (3s)
	go APIFake("API de Pagamentos", 500*time.Millisecond, canalAPI1)
	go APIFake("API de Frete", 3000*time.Millisecond, canalAPI2)

	// Limite máximo de espera: 1.5 segundos
	timeout := time.After(1500 * time.Millisecond)

	// Select captura o primeiro caso a acontecer
	select {
	case resposta := <-canalAPI1:
		fmt.Println("Vitória!", resposta)

	case resposta := <-canalAPI2:
		fmt.Println("Vitória!", resposta)

	// Se passar os 1.5s sem os canais receberem dados, roda:
	case <-timeout:
		fmt.Println("Erro: Timeout da API! O servidor demorou muito para responder.")
	}
}

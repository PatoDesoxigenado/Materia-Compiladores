/*
Exemplo 4: Concorrência Avançada (Canais e Select)
Arquivo: 04_canais_select.go
Tópicos: Channels (comunicação segura), buffers e o comando select.
Ideia de Implementação: Um simulador de "Timeout de API", fazendo duas requisições fake em paralelo e usando o
select para pegar a resposta da mais rápida ou disparar um erro de timeout se demorar muito.

ChatGPT foi ultilizado, na interface de conversação sem mais configurações, para gerar os cenários de exemplo.
Código base criado para entendimento:
package main

import (
    "fmt"
    "time"
)

func APIFake(nome string, tempoDeResposta time.Duration, canal chan<- string) {
    time.Sleep(tempoDeResposta) //simula a latência de rede
    canal <- fmt.Sprintf("Resposta bem-sucedida da %s", nome)
}

func main() {
    canalAPI1 := make(chan string)
    canalAPI2 := make(chan string)

    fmt.Println("Iniciando requisições em paralelo")
    go APIFake("API de Pagamentos", 500\*time.Millisecond, canalAPI1)
    go APIFake("API de Frete", 3000\*time.Millisecond, canalAPI2)

    timeout := time.After(1500 \* time.Millisecond)

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

Prompt enviado junto ao código: Elaborei o código acima como a implementação da ideia e dos requisitos abaixo,
adapte o código criando exemplos relevantes como cenários para os possíveis casos (API de pagamentos primeiro,
API de fretes primeiro, timeout e um exemplo que explique a difereça que buffer faz) em que essa implementação
funciona e que façam sentido segundo os tópicos.
- Tópicos: Channels (comunicação segura), buffers e o comando select.
- Contexto: Implementação de um simulador de "Timeout de API", fazendo duas requisições fake em paralelo e
usando o select para pegar a resposta da mais rápida ou disparar um erro de timeout se demorar muito.
*/

package main

import (
	"fmt"
	"time"
)

// APIFake simula uma requisição a um servidor externo.
// O canal recebe a resposta quando a simulação termina.
func APIFake(nome string, tempoDeResposta time.Duration, canal chan<- string) {
	time.Sleep(tempoDeResposta)
	canal <- fmt.Sprintf("Resposta da %s", nome)
}

// Executa um cenário com duas APIs concorrentes e um timeout.
func executarCenario(
	nome string,
	tempoPagamentos time.Duration,
	tempoFrete time.Duration,
	limite time.Duration,
) {
	fmt.Printf("\n=== %s ===\n", nome)

	// Canais com buffer de capacidade 1.
	// Cada API pode enviar sua resposta sem depender de um recebimento imediato.
	canalPagamentos := make(chan string, 1)
	canalFrete := make(chan string, 1)

	go APIFake("API de Pagamentos", tempoPagamentos, canalPagamentos)
	go APIFake("API de Frete", tempoFrete, canalFrete)

	timeout := time.After(limite)

	// Aguarda a primeira resposta ou o timeout.
	select {
	case resposta := <-canalPagamentos:
		fmt.Println("Sucesso:", resposta)

	case resposta := <-canalFrete:
		fmt.Println("Sucesso:", resposta)

	case <-timeout:
		fmt.Println("Erro: timeout! Nenhuma resposta chegou dentro do prazo.")
	}
}

// Demonstra a diferença entre canais sem buffer e com buffer.
func demonstrarBuffer() {
	fmt.Println("\n=== Cenário 4 - Canais com e sem buffer ===")

	// Canal sem buffer: o envio depende de um receptor.
	fmt.Println("\nCanal sem buffer:")
	canalSemBuffer := make(chan string)

	go func() {
		fmt.Println("API: tentando enviar a resposta...")
		canalSemBuffer <- "Resposta recebida pelo canal sem buffer"
		fmt.Println("API: envio concluído.")
	}()

	time.Sleep(200 * time.Millisecond)
	fmt.Println("Main: ainda não recebeu a resposta; o envio está bloqueado.")

	resposta := <-canalSemBuffer
	fmt.Println("Main:", resposta)

	// Canal com buffer: o envio pode ocorrer antes do recebimento.
	fmt.Println("\nCanal com buffer:")
	canalComBuffer := make(chan string, 1)

	go func() {
		fmt.Println("API: enviando a resposta...")
		canalComBuffer <- "Resposta armazenada no buffer"
		fmt.Println("API: envio concluído sem esperar pelo receptor.")
	}()

	time.Sleep(200 * time.Millisecond)

	resposta = <-canalComBuffer
	fmt.Println("Main:", resposta)
}

func main() {
	fmt.Println("Simulador de requisições concorrentes")

	// Cenário 1: Pagamentos responde primeiro.
	executarCenario(
		"Cenário 1 - Pagamentos primeiro",
		500*time.Millisecond,
		2*time.Second,
		1500*time.Millisecond,
	)

	// Cenário 2: Frete responde primeiro.
	executarCenario(
		"Cenário 2 - Frete primeiro",
		1*time.Second,
		400*time.Millisecond,
		1500*time.Millisecond,
	)

	// Cenário 3: as duas APIs demoram mais que o limite.
	executarCenario(
		"Cenário 3 - Timeout",
		2*time.Second,
		3*time.Second,
		1500*time.Millisecond,
	)

	// Cenário 4: compara o comportamento dos canais.
	demonstrarBuffer()
}

/*
Exemplo 2: Concorrência Básica (Goroutines)
- Arquivo: `02_goroutines_waitgroup.go`
- Integrante: Alerrandro Martins
- Tópicos: Goroutines e `sync.WaitGroup`.
- Ideia de Implementação: Simular um "Disparador de E-mails em Massa", onde cada envio é feito paralelamente em
    uma goroutine, usando o WaitGroup para aguardar o fim de todos os envios antes de fechar o programa.

Claude Code foi utilizado, na interface de conversação sem mais configurações, para propor os cenários
demonstrativos e auxiliar na escrita do código.

Para ver os avisos de condição de corrida que o Go detecta automaticamente:
    go run -race 02_goroutines_waitgroup.go
*/

package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// motivoInvalido identifica a falha permanente: repetir o envio não resolve.
const motivoInvalido = "endereço inválido"

// ===========================================================================
// Tipos do domínio
// ===========================================================================

// Destinatario é uma pessoa que receberá uma das mensagens do disparo.
type Destinatario struct {
	Nome  string
	Email string
}

// ResultadoEnvio registra o que aconteceu com um único destinatário.
type ResultadoEnvio struct {
	Nome       string
	Email      string
	Sucesso    bool
	Tentativas int
	Duracao    time.Duration
	Motivo     string
}

// ===========================================================================
// Seção 1 - Por que o WaitGroup é necessário
// ===========================================================================

// semWaitGroup mostra o problema: em Go, quando main termina, TODAS as goroutines
// são encerradas junto com o processo. Sem espera, os envios são cortados no meio.
func semWaitGroup() {
	fmt.Println("=== 1. O problema: disparar sem esperar ===")

	inicio := time.Now()

	for i := 1; i <= 3; i++ {
		// `go` inicia uma goroutine: uma função executada de forma concorrente,
		// com pilha própria e muito mais barata que uma thread do sistema.
		go func(numero int) {
			time.Sleep(150 * time.Millisecond) // simula a latência do envio
			fmt.Printf("  [goroutine %d] e-mail enviado!\n", numero)
		}(i)
	}

	fmt.Printf("  main terminou em %v sem esperar nenhuma goroutine.\n", time.Since(inicio))
	fmt.Println("  (as goroutines que ainda estavam trabalhando foram encerradas com o processo)")
}

// ===========================================================================
// Seção 2 - O disparo com sync.WaitGroup
// ===========================================================================

// listaDeDestinatarios é a "base" da campanha.
func listaDeDestinatarios() []Destinatario {
	return []Destinatario{
		{Nome: "Alerrandro Martins", Email: "alerrandro@uern.br"},
		{Nome: "Ana Kelry", Email: "ana.kelry@uern.br"},
		{Nome: "Jean Vitor", Email: "jean.vitor@uern.br"},
		{Nome: "Luiz Fernando", Email: "luiz.fernando@uern.br"},
		{Nome: "Thafine Jennifer", Email: "thafine@uern.br"},
		{Nome: "Maria Silva", Email: "maria.silva@gmail.com"},
		{Nome: "João Souza", Email: "joao.souza@yahoo.com"},
		{Nome: "Empresa X", Email: "contato.empresa.com"}, // sem "@": falha permanente
		{Nome: "Carlos Lima", Email: "carlos.lima@hotmail.com"},
		{Nome: "Ana Paula", Email: "ana.paula@uern.br"},
	}
}

// validaEmail imita a validação de endereço feita pelo servidor de e-mail:
// exige um "@" com texto antes e depois dele.
func validaEmail(email string) bool {
	arroba := strings.Index(email, "@")
	return arroba > 0 && arroba < len(email)-1
}

// enviarSimulado simula uma requisição ao provedor de e-mail.
func enviarSimulado(destinatario Destinatario, tentativa int) (bool, time.Duration, string) {
	// A latência é calculada a partir do e-mail, para que a demonstração
	// seja reproduzível: rodar o programa várias vezes dá o mesmo resultado.
	latencia := time.Duration(60+len(destinatario.Email)*7) * time.Millisecond
	time.Sleep(latencia)

	switch {
	case !validaEmail(destinatario.Email):
		// E-mail inválido: nenhuma tentativa nova vai resolver.
		return false, latencia, "endereço inválido"
	case tentativa == 1 && strings.HasSuffix(destinatario.Email, "@gmail.com"):
		// Falha temporária: na segunda tentativa o servidor responde bem.
		return false, latencia, "servidor temporariamente indisponível"
	default:
		return true, latencia, "entregue na caixa de entrada"
	}
}

// enviarComRetentativas tenta entregar a mensagem até `maxTentativas` vezes.
func enviarComRetentativas(destinatario Destinatario, maxTentativas int) ResultadoEnvio {
	inicio := time.Now()
	resultado := ResultadoEnvio{Nome: destinatario.Nome, Email: destinatario.Email}

	for tentativa := 1; tentativa <= maxTentativas; tentativa++ {
		sucesso, latencia, motivo := enviarSimulado(destinatario, tentativa)
		resultado.Tentativas = tentativa
		resultado.Duracao += latencia
		resultado.Motivo = motivo

		if sucesso {
			resultado.Sucesso = true
			resultado.Duracao = time.Since(inicio)
			return resultado
		}

		// Falha permanente: não adianta insistir, então encerra as tentativas.
		if motivo == motivoInvalido {
			break
		}
	}

	resultado.Duracao = time.Since(inicio)
	return resultado
}

// dispararEmMassa é o coração do exemplo: uma goroutine por destinatário,
// com o sync.WaitGroup garantindo que main só continue depois que todas terminarem.
func dispararEmMassa(destinatarios []Destinatario, maxTentativas int) []ResultadoEnvio {
	var (
		// WaitGroup funciona como um contador de tarefas pendentes.
		wg sync.WaitGroup
		// Mutex protege o slice de resultados, que é escrito por várias
		// goroutines ao mesmo tempo (estado compartilhado).
		mu sync.Mutex
		// atomic.Int64 é um contador seguro para atualizar sem travar.
		enviados atomic.Int64
	)

	resultados := make([]ResultadoEnvio, len(destinatarios))

	// Add deve ser chamada ANTES de `go`, para que a contagem não fique
	// zerada enquanto as goroutines ainda estão começando.
	wg.Add(len(destinatarios))

	for i, destinatario := range destinatarios {
		go func(indice int, d Destinatario) {
			// Done é adiado para a saída da goroutine: assim ele roda
			// exatamente uma vez, mesmo se o código sair por um return.
			defer wg.Done()

			resultado := enviarComRetentativas(d, maxTentativas)
			resultados[indice] = resultado

			// Como cada goroutine escreve em uma posição diferente do slice,
			// a trava só é necessária no trecho de leitura do contador/log.
			mu.Lock()
			concluidos := enviados.Load() + 1
			enviados.Store(concluidos)
			fmt.Printf("  [%d/%d] %-18s %s\n", concluidos, len(destinatarios),
				d.Nome, statusDe(resultado))
			mu.Unlock()
		}(i, destinatario)
	}

	// Wait bloqueia a goroutine atual (a main) até o contador chegar a zero.
	wg.Wait()

	return resultados
}

// statusDe resume o resultado em uma frase curta.
func statusDe(r ResultadoEnvio) string {
	if r.Sucesso {
		return fmt.Sprintf("OK em %d tentativa(s) (%v)", r.Tentativas, r.Duracao.Round(time.Millisecond))
	}
	return fmt.Sprintf("FALHOU após %d tentativa(s): %s", r.Tentativas, r.Motivo)
}

// ===========================================================================
// Seção 3 - Concorrência limitada por lotes
// ===========================================================================

// dispararEmLotes evita criar centenas de goroutines de uma vez: cada lote roda
// com, no máximo, `tamanhoDoLote` envios em paralelo e o próximo lote espera.
func dispararEmLotes(destinatarios []Destinatario, tamanhoDoLote, maxTentativas int) []ResultadoEnvio {
	var (
		wg         sync.WaitGroup
		mu         sync.Mutex
		enviados   atomic.Int64
		resultados []ResultadoEnvio
	)

	inicio := time.Now()

	for inicioLote := 0; inicioLote < len(destinatarios); inicioLote += tamanhoDoLote {
		fimLote := inicioLote + tamanhoDoLote
		if fimLote > len(destinatarios) {
			fimLote = len(destinatarios)
		}
		lote := destinatarios[inicioLote:fimLote]

		wg.Add(len(lote))

		for _, destinatario := range lote {
			go func(d Destinatario) {
				defer wg.Done()

				resultado := enviarComRetentativas(d, maxTentativas)

				mu.Lock()
				resultados = append(resultados, resultado)
				concluidos := enviados.Add(1)
				fmt.Printf("  [%d/%d] %-18s %s\n", concluidos, len(destinatarios),
					d.Nome, statusDe(resultado))
				mu.Unlock()
			}(destinatario)
		}

		wg.Wait() // só depois de Wait o próximo lote é disparado
	}

	fmt.Printf("  %d lotes de até %d concluídos em %v\n",
		(len(destinatarios)+tamanhoDoLote-1)/tamanhoDoLote, tamanhoDoLote, time.Since(inicio).Round(time.Millisecond))

	return resultados
}

// ===========================================================================
// Seção 4 - Sequencial x concorrente
// ===========================================================================

// compararDesempenho mostra o ganho prático: o mesmo lote executado de forma
// sequencial e depois em paralelo.
func compararDesempenho(destinatarios []Destinatario, maxTentativas int) {
	fmt.Println("\n=== 4. Sequencial x concorrente (mesmo lote) ===")

	inicio := time.Now()
	for _, d := range destinatarios {
		enviarComRetentativas(d, maxTentativas)
	}
	sequencial := time.Since(inicio)

	inicio = time.Now()
	dispararEmMassa(destinatarios, maxTentativas)
	concorrente := time.Since(inicio)

	fmt.Printf("  Sequencial:  %v\n", sequencial.Round(time.Millisecond))
	fmt.Printf("  Concorrente: %v\n", concorrente.Round(time.Millisecond))
	if concorrente > 0 {
		fmt.Printf("  Ganho:       %.2fx mais rápido\n", float64(sequencial)/float64(concorrente))
	}
}

// ===========================================================================
// Relatório final
// ===========================================================================

// imprimirRelatorio resume os resultados. Como as goroutines terminam em ordem
// aleatória, ordenamos por e-mail para deixar a saída determinística.
func imprimirRelatorio(resultados []ResultadoEnvio) {
	ordenados := make([]ResultadoEnvio, len(resultados))
	copy(ordenados, resultados)
	sort.Slice(ordenados, func(i, j int) bool {
		return ordenados[i].Email < ordenados[j].Email
	})

	var (
		enviados              int
		falhas                int
		tentativas            int
		sucessoEmUmaTentativa int
	)

	fmt.Println("\n--- Relatório da campanha ---")
	for _, r := range ordenados {
		fmt.Printf("  %-18s %-26s %-8s %v\n", r.Nome, r.Email, marcaDe(r.Sucesso), statusDe(r))
		tentativas += r.Tentativas
		if r.Tentativas == 1 && r.Sucesso {
			sucessoEmUmaTentativa++
		}
		if r.Sucesso {
			enviados++
		} else {
			falhas++
		}
	}

	fmt.Println()
	fmt.Printf("  Total de destinatários: %d\n", len(ordenados))
	fmt.Printf("  Enviados com sucesso:   %d\n", enviados)
	fmt.Printf("  Falhas:                 %d\n", falhas)
	fmt.Printf("  Tentativas de envio:    %d\n", tentativas)
	fmt.Printf("  Sucesso na 1ª tentativa: %d\n", sucessoEmUmaTentativa)
}

func marcaDe(sucesso bool) string {
	if sucesso {
		return "OK"
	}
	return "FALHOU"
}

func main() {
	fmt.Println("=== EXEMPLO 2: GOROUTINES E sync.WaitGroup ===")
	fmt.Println("Disparador de E-mails em Massa")
	fmt.Println()

	semWaitGroup()

	const maxTentativas = 3

	fmt.Println("\n=== 2. Disparo com WaitGroup (todos em paralelo) ===")
	destinatarios := listaDeDestinatarios()
	inicio := time.Now()
	resultados := dispararEmMassa(destinatarios, maxTentativas)
	fmt.Printf("  campanha concluída em %v\n", time.Since(inicio).Round(time.Millisecond))

	fmt.Println("\n=== 3. Disparo em lotes de 3 (concorrência limitada) ===")
	dispararEmLotes(destinatarios, 3, maxTentativas)

	compararDesempenho(destinatarios, maxTentativas)

	imprimirRelatorio(resultados)
}

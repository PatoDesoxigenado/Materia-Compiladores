# Integrante: Alerrandro Martins

Os dois arquivos desta pasta são exemplos da **linguagem de origem (Go)** e não recebem
nenhum argumento pela linha de comando: basta compilá-los e executá-los.

---

## `01_tipos_variaveis.go` — Tipos de Dados e Variáveis

> Tópicos: tipos primitivos, compostos (`struct`, `slice`, `map`), escopo e *zero values*.

A ideia é um **Carrinho de Compras**. O exemplo está dividido em seis seções que vão do
mais simples ao mais composto:

| Seção | Conteúdo |
|---|---|
| 1 | Tipos primitivos: tipo estático, valor zero e tamanho em memória (`unsafe.Sizeof`) |
| 2 | *Zero values*: `var x T` nunca fica com lixo de memória; slice/map `nil`, leitura de map `nil` |
| 3 | *Structs*: literais nomeados vs. posicional, aninhamento e cópia por valor |
| 4 | *Slice*: `len` x `cap`, `append`, `copy`, fatiamento e o *aliasing* do array subjacente |
| 5 | *Escopo*: pacote, função e bloco, com sombreamento (`:=` vs. `=`) |
| 6 | O carrinho montado: filtro do catálogo, subtotal, cupom e agrupamento por categoria (`map`) |

Alguns pontos que o exemplo deixa visíveis:

- `string` ocupa **2 palavras** (endereço dos dados + comprimento), não 1 byte por caractere;
- um `rune` (`Ç`) ocupa **4 bytes**, enquanto o texto `"compiladores"` ocupa **12**;
- `*Produto` ocupa **1 palavra**, porque guarda o endereço da struct e não a struct inteira;
- fatiamentos (`s[i:j]`) **compartilham o array subjacente**: alterar `s[1:3]` altera a original;
- `float64` é binário de precisão, então `0.1 + 0.2` não resulta em `0.3` exatamente.

---

## `02_goroutines_waitgroup.go` — Goroutines e `sync.WaitGroup`

> Tópicos: goroutines e `sync.WaitGroup`.

A ideia é um **Disparador de E-mails em Massa**, com quatro cenários:

1. **O problema** — sem `WaitGroup`, `main` termina e o processo encerra as goroutines no meio do envio;
2. **Disparo com `WaitGroup`** — uma goroutine por destinatário, `wg.Add(...)` antes do `go`,
   `defer wg.Done()`, `sync.Mutex` protegendo o slice de resultados e `sync/atomic` no contador;
3. **Lotes de 3** — `wg.Wait()` entre os lotes para limitar quantas goroutines rodam ao mesmo tempo;
4. **Sequencial x concorrente** — a mesma lista de envios medida nos dois modelos.

A simulação inclui **tentativas automáticas**: um endereço `@gmail.com` falha na primeira
tentativa e tem sucesso na segunda (falha temporária), enquanto um endereço sem `@` é uma
falha permanente — o laço de retentativas encerra imediatamente nesse caso.

> Canais e `select` **não** são usados aqui, porque esses tópicos já são do exemplo 04
> (Thafine Jennifer).

---

## Como compilar e executar

Cada arquivo tem que ser compilado individualmente (ambos definam `func main`):

```
go build 01_tipos_variaveis.go
go build 02_goroutines_waitgroup.go
```

E depois executado diretamente:

```
Exemplos no Ubuntu:
  $ ./01_tipos_variaveis
  $ ./02_goroutines_waitgroup
```

Ou, para testar sem gerar binário:

```
go run 01_tipos_variaveis.go
go run 02_goroutines_waitgroup.go
```

### Detector de condições de corrida

O exemplo 02 compartilha estado entre goroutines. Para deixar o Go auditar esse
acesso, execute com a flag `-race`:

```
go run -race 02_goroutines_waitgroup.go
```

Sem avisos no terminal = o slice de resultados está sendo protegido corretamente.

---

## Observações sobre a saída

A saída completa e literal dos dois exemplos está em [`resultado.txt`](resultado.txt).

Os tempos de execução **variam** de uma execução para outra, porque as latências são
simuladas com `time.Sleep`. Já a *ordem* das linhas é reprodutível: a latência de cada
destinatário é calculada a partir do e-mail, e o relatório final é ordenado por e-mail
(a ordem de término das goroutines é, por natureza, imprevisível).

---

#### Uso de IA generativa

Ambos os exemplos foram desenvolvidos com o apoio do **Claude Code** (interface de
conversação, sem configurações adicionais), usado para propor os cenários demonstrativos
e auxiliar na escrita do código. O código foi revisado, formatado com `gofmt`, validado
com `go vet` e testado com o detector de condições de corrida do Go (`-race`).

# Matéria-Compiladores

Repositório da Disciplina COMPILADORES E PARADIGMAS DE PROGRAMAÇÃO - UERN Semestre 2026.2


## 📂 Estrutura de Pastas

```
Matéria-Compiladores/
├── LICENSE
├── README.md
└── exemplos-origem/
    ├── README.md
    ├── Alerrandro Martins/
    │   ├── 01_tipos_variaveis.go
    │   └── 02_goroutines_waitgroup.go
    ├── Thafine Jennifer/
    │   ├── 03_ponteiros_referencias.go
    │   └── 04_canais_select.go
    ├── Ana Kelry/
    │   ├── 05_condicionais_switch.go
    │   └── 06_tratamento_erros.go
    ├── Jean Vitor/
    │   ├── 07_loops_defer.go
    │   └── 08_funcoes.go
    └── Luiz Fernando/
        ├── 09_closures_lambdas.go
        └── 10_generics.go
```

## 📁 Conteúdo das Pastas

* **`exemplos-origem/`**: Contém 10 códigos-fonte escritos na linguagem de origem (Go), criados pela equipe para demonstrar os aspectos léxicos, sintáticos e semânticos da linguagem, cobrindo desde tipos básicos até concorrência nativa. Os arquivos estão organizados em subpastas por integrante responsável.

## 👥 Divisão de Tarefas (10 Exemplos)

Cada integrante do grupo ficou responsável por codificar, comentar e testar **2 exemplos distintos**, garantindo que não haja repetição de temas.

### 1️⃣ Integrante: Alerrandro Martins
* **Exemplo 1: Tipos de Dados e Variáveis**
  * **Arquivo:** `01_tipos_variaveis.go`
  * **Tópicos:** Tipos primitivos, compostos (structs, slices, maps), escopo e *zero values*.
  * **Ideia de Implementação:** Criar um sistema simples de "Carrinho de Compras" usando `struct` para os produtos, um `slice` para a lista e um `map` para categorizar os itens.
* **Exemplo 2: Concorrência Básica (Goroutines)**
  * **Arquivo:** `02_goroutines_waitgroup.go`
  * **Tópicos:** Goroutines e `sync.WaitGroup`.
  * **Ideia de Implementação:** Simular um "Disparador de E-mails em Massa", onde cada envio é feito paralelamente em uma goroutine, usando o WaitGroup para aguardar o fim de todos os envios antes de fechar o programa.

### 2️⃣ Integrante: Thafine Jennifer
* **Exemplo 3: Ponteiros e Passagem de Parâmetros**
  * **Arquivo:** `03_ponteiros_referencias.go`
  * **Tópicos:** Passagem por valor vs. simulação de referência usando ponteiros explícitos (`*T` e `&x`).
  * **Ideia de Implementação:** Criar um sistema de "Transação Bancária", demonstrando que tentar alterar o saldo passando a struct por valor não funciona, exigindo o uso de ponteiros para efetivar o depósito/saque.
* **Exemplo 4: Concorrência Avançada (Canais e Select)**
  * **Arquivo:** `04_canais_select.go`
  * **Tópicos:** Channels (comunicação segura), buffers e o comando `select`.
  * **Ideia de Implementação:** Um simulador de "Timeout de API", fazendo duas requisições fake em paralelo e usando o `select` para pegar a resposta da mais rápida ou disparar um erro de timeout se demorar muito.

### 3️⃣ Integrante: Ana Kelry
* **Exemplo 5: Condicionais (If e Switch)**
  * **Arquivo:** `05_condicionais_switch.go`
  * **Tópicos:** `if` com inicialização curta, `switch` (sem break implícito) e `type switch`.
  * **Ideia de Implementação:** Um "Analisador de Permissões de Usuário" que verifica os papéis (Admin, User, Guest) com um `switch`, e usa o `type switch` para descobrir dinamicamente se o payload recebido é uma string ou um int.
* **Exemplo 6: Tratamento de Erros (Error é Valor)**
  * **Arquivo:** `06_tratamento_erros.go`
  * **Tópicos:** Interface `error`, formatação com `%w` e inspeção com `errors.Is` ou `errors.As`.
  * **Ideia de Implementação:** Um "Validador de Cadastro". Criar uma struct de erro customizada (`ErroValidacao`) que informa qual campo falhou e o motivo, testando o erro no fluxo principal.

### 4️⃣ Integrante: Jean Vitor
* **Exemplo 7: Repetições e Desvios (For e Defer)**
  * **Arquivo:** `07_loops_defer.go`
  * **Tópicos:** O único laço `for` do Go (clássico, condicional e range) e a cláusula `defer`.
  * **Ideia de Implementação:** Um simulador de "Leitura e Processamento de Logs". Usar o `range` para iterar sobre linhas de log, e o `defer` para garantir o fechamento seguro do arquivo simulado no final do processo, independentemente de erros.
* **Exemplo 8: Funções (Retornos e Variádicos)**
  * **Arquivo:** `08_funcoes.go`
  * **Tópicos:** Múltiplos retornos, retornos nomeados e parâmetros variádicos.
  * **Ideia de Implementação:** Uma função "Calculadora de Estatísticas" que recebe uma quantidade infinita de notas (variádico), e retorna três valores: o total, a média e um erro (caso nenhuma nota seja enviada).

### 5️⃣ Integrante: Luiz Fernando
* **Exemplo 9: Subprogramas Delegados (Closures e Lambdas)**
  * **Arquivo:** `09_closures_lambdas.go`
  * **Tópicos:** Funções como cidadãs de primeira classe e retenção de estado léxico (closures).
  * **Ideia de Implementação:** Criar um "Gerador Sequencial de IDs" ou de "Tickets" usando um *closure* que encapsula e mantém o estado da variável contadora vivo entre as chamadas da função.
* **Exemplo 10: Programação Genérica (Generics)**
  * **Arquivo:** `10_generics.go`
  * **Tópicos:** Uso de Type Parameters `[T any]` introduzidos no Go 1.18+.
  * **Ideia de Implementação:** Criar uma "Estrutura de Pilha (Stack) Genérica". A mesma estrutura deve ser capaz de empilhar e desempilhar valores `int`, `string` ou `structs`, demonstrando a reutilização de código sem sacrificar a tipagem estática.

  ## 🤖 Uso de Ferramentas de IA Generativa

Durante o desenvolvimento dos códigos de exemplo desta pasta, os integrantes do grupo utilizaram ferramentas de Inteligência Artificial generativa (variando entre **Claude/Claude Code**, **Gemini** e **ChatGPT**) como apoio ao aprendizado.

Como a construção foi colaborativa e iterativa entre os 5 membros, o uso e as configurações das ferramentas variaram, focando principalmente em prompts de assistência para as seguintes tarefas:
* **Geração inicial:** Ideias de cenários práticos e criação da estrutura base dos códigos em Go.
* **Refatoração:** Sugestões para adequar os exemplos às boas práticas e ao estilo de escrita da linguagem (Go *idiomatic*).
* **Correção e Complemento:** Auxílio na identificação de erros de sintaxe/compilação e complemento de trechos de código.

As IAs foram utilizadas estritamente como assistentes de programação, garantindo que os exemplos criados refletissem corretamente o funcionamento da linguagem estudada.

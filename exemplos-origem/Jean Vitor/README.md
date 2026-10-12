### Informações:

Ambos os arquivos `07_loops_defer.go` e `08_funcoes.go` são executados recebendo entradas através do terminal.

---

O arquivo `07_loops_defer.go` tem como objetivo exemplificar o uso do único laço do Go: o laço `for` com suas variações (clássico, condicional e range), e a cláusula `defer`.

>A ideia de implementação é um simples simulador de leitura e logs, recebendo como parâmetro um número inteiro, que representa a quantidade de linhas que se deseja visualizar (começando pelo mais recente). Depois, pede a localização de um log, como o `sample.log`.

O arquivo `08_funcoes.go` tem como objetivo exemplificar funções, múltiplos retornos, retornos nomeados e parâmetros variádicos.

> A ideia de implementação é uma calculadora de estatísticas fraca das pernas que recebe uma quantidade infinita de notas (variádico), e retorna o total e a média, ou um erro, caso nenhuma nota seja enviada. No terminal, recebe inteiros ou floats separados por espaços.

---

Cada arquivo tem que ser passado por `go build`:
```
go build 07_loops_defer.go
go build 08_funcoes.go
```
E depois pode ser executado diretamente:
```
Exemplos no ubuntu:
  $ ./07_loops_defer 1
    | Insira o local do arquivo .log: sample.log
    | ==========
    | 2026-11-10 13:20:13 - Successful: user registration.
    | ==========

  $ ./08_funcoes 9 10 8 8.5 9.2 10
    | Somatório total das notas:  54.7
    | Média das notas:  9.116666666666667
```

---
#### Não foi feito o uso de IAs nestes exemplos

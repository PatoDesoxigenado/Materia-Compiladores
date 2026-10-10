# Exemplos de Código - Linguagem Go

Esta pasta contém os 10 arquivos de códigos-fonte de exemplo na linguagem de origem (Go), distribuídos entre os integrantes da equipe. 

## 🛠️ Compilador Utilizado
O compilador utilizado para rodar estes exemplos é o compilador oficial da linguagem Go, conhecido como **`gc` (Go Compiler)**, que já vem embutido por padrão na *toolchain* oficial disponibilizada pelo Google.

## 📥 Como Instalar o Compilador
Para compilar e executar os códigos, você precisa ter a linguagem Go instalada na sua máquina.

1. Acesse o site oficial: [https://go.dev/dl/](https://go.dev/dl/)
2. Baixe o instalador correspondente ao seu sistema operacional (Windows, Linux ou macOS).
3. Execute o instalador e siga o assistente padrão (*Next > Install*).
4. Para verificar se a instalação foi bem-sucedida, abra um novo terminal (Prompt de Comando, PowerShell ou bash) e digite:
   ```bash```
   go version

## ▶️ Como Compilar e Executar os Arquivos Fontes

A ferramenta de linha de comando do Go permite duas abordagens principais para rodar os exemplos desta pasta. Abra o terminal na raiz desta pasta e escolha um dos métodos abaixo:

### Método 1: Execução Direta (Recomendado para testes)
O comando `go run` compila o código em um diretório temporário e o executa imediatamente. É a forma mais rápida de testar os scripts.
```bash```
go run nome_do_arquivo.go

### Método 2: Compilar e depois Executar
Se você quiser gerar o arquivo binário (executável final) a partir do código-fonte:
1. Compile o arquivo com o comando `go build`:
   ```bash```
   go build nome_do_arquivo.go
2. Isso vai gerar um executável nativo na mesma pasta (ex: `nome_do_arquivo.exe` no Windows ou `./nome_do_arquivo` no Linux). Execute o binário gerado:
   ```bash```
   ./nome_do_arquivo
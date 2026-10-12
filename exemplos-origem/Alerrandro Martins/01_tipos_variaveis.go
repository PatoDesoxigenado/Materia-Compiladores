/*
Exemplo 1: Tipos de Dados e Variáveis
- Arquivo: `01_tipos_variaveis.go`
- Integrante: Alerrandro Martins
- Tópicos: Tipos primitivos, compostos (structs, slices, maps), escopo e *zero values*.
- Ideia de Implementação: Criar um sistema simples de "Carrinho de Compras" usando `struct` para os produtos,
    um `slice` para a lista de itens e um `map` para categorizar os itens.

Claude Code foi utilizado, na interface de conversação sem mais configurações, para propor os cenários
demonstrativos (seções 1 a 5) e auxiliar na escrita do código.
*/

package main

import (
	"fmt"
	"math"
	"sort"
	"unsafe"
)

// ===========================================================================
// Tipos compostos do domínio: as structs que descrevem o carrinho
// ===========================================================================

// Endereco é uma struct simples que aparece como campo de Cliente (struct aninhada).
type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	CEP        string
}

// Cliente agrupa os dados de quem está comprando.
type Cliente struct {
	Nome             string
	Email            string
	Endereco         Endereco // struct dentro de struct
	AssinaNewsletter bool
}

// Produto é o registro de um item do catálogo.
// Os campos são exportados (maiúscula inicial) porque o exemplo simula um módulo
// de catálogo que seria consumido por outros pacotes.
type Produto struct {
	Codigo    int
	Nome      string
	Preco     float64
	Estoque   int
	Ativo     bool
	Categoria string
}

// cupomPromocional tem escopo de pacote: existe durante toda a execução do programa.
const cupomPromocional = 0.15

// ===========================================================================
// Seção 1 - Tipos primitivos
// ===========================================================================

// demoPrimitivos mostra que, em Go, todo valor tem um tipo estático definido em
// tempo de compilação e um tamanho fixo em memória.
func demoPrimitivos() {
	fmt.Println("=== 1. Tipos primitivos: tipo estático, valor zero e tamanho ===")

	// Declaração agrupada: cada linha declara uma variável já inicializada.
	var (
		inteiro  int      = 42
		pequeno  int8     = 127
		semSinal uint     = 42
		real     float64  = 3.14
		texto    string   = "compiladores"
		logico   bool     = true
		letra    rune     = 'Ç' // rune = int32: representa um ponto de código Unicode
		bruto    byte     = 65  // byte  = uint8: representa um único byte ('A')
		ponteiro *Produto       // guarda um endereço de memória, não o valor apontado
	)

	// %T imprime o tipo estático; %v o valor; %q o valor como string entre aspas.
	// Os textos já vêm formatados para que string e rune apareçam entre aspas.
	linhas := []struct {
		tipo     string
		valorFmt string
		zeroFmt  string
		tamanho  uintptr
	}{
		{"int", fmt.Sprintf("%d", inteiro), "0", unsafe.Sizeof(inteiro)},
		{"int8", fmt.Sprintf("%d", pequeno), "0", unsafe.Sizeof(pequeno)},
		{"uint", fmt.Sprintf("%d", semSinal), "0", unsafe.Sizeof(semSinal)},
		{"float64", fmt.Sprintf("%v", real), "0", unsafe.Sizeof(real)},
		{"string", fmt.Sprintf("%q", texto), `""`, unsafe.Sizeof(texto)},
		{"bool", fmt.Sprintf("%v", logico), "false", unsafe.Sizeof(logico)},
		{"rune (int32)", fmt.Sprintf("%q (%d)", letra, letra), "0", unsafe.Sizeof(letra)},
		{"byte (uint8)", fmt.Sprintf("%d (%q)", bruto, bruto), "0", unsafe.Sizeof(bruto)},
		{"*Produto", fmt.Sprintf("%v", ponteiro), "<nil>", unsafe.Sizeof(ponteiro)},
	}

	fmt.Printf("  %-14s %-20s %-10s %s\n", "TIPO", "VALOR", "ZERO VALUE", "TAMANHO")
	for _, l := range linhas {
		fmt.Printf("  %-14s %-20s %-10s %d byte(s)\n", l.tipo, l.valorFmt, l.zeroFmt, l.tamanho)
	}

	fmt.Println("  Observações:")
	fmt.Println("    - string ocupa 2 palavras (endereço dos dados + comprimento), não 1 byte por caractere.")
	fmt.Println("    - *Produto ocupa 1 palavra: ele guarda o endereço da struct, não a struct inteira.")
	fmt.Printf("    - a letra %q é um rune (código %d) e ocupa %d bytes, mas o texto completo %q ocupa %d bytes.\n",
		letra, letra, unsafe.Sizeof(letra), "compiladores", len("compiladores"))

	// Atenção semântica: float64 representa números em base 2, então 0.1 não é exato.
	fmt.Println("    - float64 é binário de precisão: 0.1+0.2 não dá 0.3 exatamente.")
	fmt.Printf("      0.1 + 0.2 = %.20f\n", 0.1+0.2)
	fmt.Printf("      arredondado para 2 casas = %.2f\n", math.Round((0.1+0.2)*100)/100)
}

// ===========================================================================
// Seção 2 - Zero values
// ===========================================================================

// demoZeroValues mostra que `var x T` nunca fica com "lixo de memória":
// o compilador garante a inicialização automática com o valor zero do tipo.
func demoZeroValues() {
	fmt.Println("\n=== 2. Zero values: `var x T` sempre nasce inicializado ===")

	var (
		quantidade int
		preco      float64
		descricao  string
		disponivel bool
		tamanhos   []string           // zero de slice = nil
		descontos  map[string]float64 // zero de map  = nil
	)

	// Uma struct inteira também nasce zerada, campo a campo.
	var cliente Cliente

	fmt.Printf("  int     -> %d\n", quantidade)
	fmt.Printf("  float64 -> %.2f\n", preco)
	fmt.Printf("  string  -> %q (len=%d)\n", descricao, len(descricao))
	fmt.Printf("  bool    -> %v\n", disponivel)
	fmt.Printf("  slice   -> %v (len=%d, cap=%d)\n", tamanhos, len(tamanhos), cap(tamanhos))
	fmt.Printf("  map     -> %v (len=%d)\n", descontos, len(descontos))
	fmt.Printf("  Cliente -> nome=%q cidade=%q newsletter=%v\n",
		cliente.Nome, cliente.Endereco.Cidade, cliente.AssinaNewsletter)

	// Ler de um map nil é seguro e devolve o valor zero do tipo do valor.
	fmt.Printf("  leitura em map nil: descontos[\"frete\"] = %.2f\n", descontos["frete"])
	fmt.Println("  (escrever em map nil causaria panic: é preciso antes usar make)")

	// append em slice nil funciona: o runtime aloca o array subjacente.
	tamanhos = append(tamanhos, "P", "M", "G")
	fmt.Printf("  após append no slice nil: %v (len=%d, cap=%d)\n", tamanhos, len(tamanhos), cap(tamanhos))
}

// ===========================================================================
// Seção 3 - Structs
// ===========================================================================

// demoStructs mostra literais compostos, campos nomeados, aninhamento e o fato de
// que uma struct é um valor (a cópia é independente do original).
func demoStructs() {
	fmt.Println("\n=== 3. Structs: literais, campos nomeados e aninhamento ===")

	cliente := Cliente{
		Nome:  "Alerrandro Martins",
		Email: "alerrandro@uern.br",
		Endereco: Endereco{ // literal aninhado, com campos nomeados
			Logradouro: "Rua das Acácias",
			Numero:     120,
			Cidade:     "Mossoró",
			CEP:        "59633-070",
		},
		AssinaNewsletter: true,
	}

	// Literal posicional: os valores seguem a ordem de declaração dos campos.
	// Compila, mas é frágil — se alguém inserir um campo novo, a ordem quebra.
	posicional := Endereco{"Av. Presidente Dutra", 45, "Parnamirim", "59633-000"}

	fmt.Printf("  Cliente: %s <%s>\n", cliente.Nome, cliente.Email)
	fmt.Printf("  Endereço aninhado:  %s, %d - %s/%s\n",
		cliente.Endereco.Logradouro, cliente.Endereco.Numero, cliente.Endereco.Cidade, cliente.Endereco.CEP)
	fmt.Printf("  Endereço posicional: %s, %d - %s/%s\n",
		posicional.Logradouro, posicional.Numero, posicional.Cidade, posicional.CEP)
	fmt.Printf("  Assina newsletter: %v\n", cliente.AssinaNewsletter)

	// Atribuição entre structs copia todos os campos.
	copia := cliente
	copia.Endereco.Numero = 999
	fmt.Printf("  original tem numero %d, mas a copia tem %d\n",
		cliente.Endereco.Numero, copia.Endereco.Numero)
	fmt.Println("  (structs são valores: a cópia é independente do original)")
}

// ===========================================================================
// Seção 4 - Slice
// ===========================================================================

// demoSlice manipula um slice: len, cap, append, copy, fatiamento e remoção.
func demoSlice() {
	fmt.Println("\n=== 4. Slice: len, cap, append, copy e fatiamento ===")

	// Slice é um descritor de 3 palavras: ponteiro para o array + len + cap.
	// Um literal já nasce com len == cap == número de elementos.
	lista := []string{"camiseta", "calca", "bone"}
	fmt.Printf("  literal: %v (len=%d, cap=%d)\n", lista, len(lista), cap(lista))

	// append pode preservar a capacidade ou realocar o array (aí o cap dobra).
	lista = append(lista, "tenis", "meia")
	fmt.Printf("  após append: %v (len=%d, cap=%d)\n", lista, len(lista), cap(lista))

	// Fatiamento [i:j] cria um novo descritor que compartilha o MESMO array.
	meio := lista[1:3]
	fmt.Printf("  lista[1:3] = %v\n", meio)
	meio[0] = "CAMISA"
	fmt.Printf("  alterando meio[0], a lista original virou %v (array compartilhado)\n", lista)

	// copy cria um array independente.
	independente := make([]string, len(lista))
	copy(independente, lista) // retorna quantos elementos foram copiados
	independente[0] = "jaqueta"
	fmt.Printf("  com copy, a lista continua %v e a cópia é %v\n", lista, independente)

	// Removendo um item: recorta o fatiamento e desloca os demais.
	lista = append(lista[:1], lista[2:]...)
	fmt.Printf("  após remover o segundo item: %v\n", lista)

	// len vs cap em fatiamento: dá para estender a visão sem alocar memória nova.
	// O array subjacente é maior que o slice, então aparece uma posição vazia.
	ampliado := lista[:cap(lista)]
	fmt.Printf("  lista[:cap] expõe o array inteiro: %v (len=%d)\n", ampliado, len(ampliado))

	// make aloca e já define len e cap.
	preVazado := make([]int, 0, 4)
	preVazado = append(preVazado, 10, 20, 30)
	fmt.Printf("  make([]int, 0, 4) + append: %v (len=%d, cap=%d)\n", preVazado, len(preVazado), cap(preVazado))
}

// ===========================================================================
// Seção 5 - Escopo
// ===========================================================================

// demoEscopo mostra os três níveis de escopo: pacote, função e bloco.
func demoEscopo() {
	fmt.Println("\n=== 5. Escopo: pacote, função e bloco ===")

	// Cupom declarado no escopo de pacote: acessível em qualquer função do arquivo.
	fmt.Printf("  escopo de pacote: cupomPromocional = %.0f%%\n", cupomPromocional*100)

	desconto := 0.05 // escopo da função demoEscopo
	fmt.Printf("  escopo da função: desconto = %.0f%%\n", desconto*100)

	{
		// `:=` cria uma NOVA variável que sombreia (shadow) a externa.
		// `=` apenas altera a variável que já existe.
		desconto := 0.30
		fmt.Printf("  dentro do bloco (sombreada): desconto = %.0f%%\n", desconto*100)
	}

	fmt.Printf("  fora do bloco, a externa continua: desconto = %.0f%%\n", desconto*100)

	// A variável declarada dentro do bloco não existe fora dele.
	// Descomente a linha abaixo para ver o erro de compilação "undefined: total":
	//
	// fmt.Printf("  total = %d\n", total)

	total := 120 // escopo da função, declarado após o bloco
	fmt.Printf("  total declarado na função: %d\n", total)
}

// ===========================================================================
// Lógica do carrinho
// ===========================================================================

// catalogoBase é o slice de produtos disponível no "banco de dados" da loja.
func catalogoBase() []Produto {
	return []Produto{
		{Codigo: 101, Nome: "Camiseta básica", Preco: 49.90, Estoque: 12, Ativo: true, Categoria: "vestuario"},
		{Codigo: 102, Nome: "Calça jeans", Preco: 129.50, Estoque: 5, Ativo: true, Categoria: "vestuario"},
		{Codigo: 103, Nome: "Tênis urbano", Preco: 259.00, Estoque: 0, Ativo: true, Categoria: "calcados"},
		{Codigo: 104, Nome: "Meia de algodão", Preco: 12.50, Estoque: 40, Ativo: true, Categoria: "vestuario"},
		{Codigo: 105, Nome: "Boné bordado", Preco: 35.00, Estoque: 8, Ativo: false, Categoria: "acessorios"},
		{Codigo: 106, Nome: "Mochila 25L", Preco: 199.90, Estoque: 3, Ativo: true, Categoria: "acessorios"},
	}
}

// agruparPorCategoria recebe um slice e devolve um map de categoria -> produtos.
// É o vínculo prático entre os dois tipos compostos do exemplo.
func agruparPorCategoria(produtos []Produto) map[string][]Produto {
	grupos := make(map[string][]Produto)

	for _, p := range produtos {
		grupos[p.Categoria] = append(grupos[p.Categoria], p)
	}

	return grupos
}

// montarCarrinho filtra o catálogo (só entra item ativo e com estoque) e calcula o subtotal.
func montarCarrinho(catalogo []Produto) (escolhidos []Produto, subtotal float64, descartados []string) {
	for _, p := range catalogo {
		if !p.Ativo {
			descartados = append(descartados, fmt.Sprintf("%s (inativo)", p.Nome))
			continue
		}
		if p.Estoque <= 0 {
			descartados = append(descartados, fmt.Sprintf("%s (sem estoque)", p.Nome))
			continue
		}

		escolhidos = append(escolhidos, p)
		subtotal += p.Preco
	}

	return escolhidos, subtotal, descartados
}

// imprimirGrupos percorre o map. A ordem de iteração de um map não é garantida,
// então ordenamos as chaves para deixar a saída determinística.
func imprimirGrupos(grupos map[string][]Produto) {
	chaves := make([]string, 0, len(grupos))
	for categoria := range grupos {
		chaves = append(chaves, categoria)
	}
	sort.Strings(chaves)

	for _, categoria := range chaves {
		itens := grupos[categoria]
		fmt.Printf("  %s (%d item(ns)):\n", categoria, len(itens))

		for _, p := range itens {
			fmt.Printf("    [%03d] %-16s R$ %7.2f  estoque: %d\n", p.Codigo, p.Nome, p.Preco, p.Estoque)
		}
	}
}

func main() {
	fmt.Println("=== EXEMPLO 1: TIPOS DE DADOS E VARIÁVEIS ===")
	fmt.Println("Sistema de Carrinho de Compras")
	fmt.Println()

	demoPrimitivos()
	demoZeroValues()
	demoStructs()
	demoSlice()
	demoEscopo()

	fmt.Println("\n=== 6. O carrinho montado ===")

	catalogo := catalogoBase()
	escolhidos, subtotal, descartados := montarCarrinho(catalogo)

	fmt.Printf("  Itens no carrinho: %d\n", len(escolhidos))
	for _, p := range escolhidos {
		fmt.Printf("    [%03d] %-16s R$ %7.2f  (%s)\n", p.Codigo, p.Nome, p.Preco, p.Categoria)
	}

	if len(descartados) > 0 {
		fmt.Println("  Itens fora do carrinho:")
		for _, motivo := range descartados {
			fmt.Printf("    - %s\n", motivo)
		}
	}

	desconto := subtotal * cupomPromocional
	total := math.Round((subtotal-desconto)*100) / 100

	fmt.Printf("\n  Subtotal:  R$ %8.2f\n", subtotal)
	fmt.Printf("  Cupom (%2.0f%%): -R$ %7.2f\n", cupomPromocional*100, desconto)
	fmt.Printf("  Total:     R$ %8.2f\n", total)

	fmt.Println("\n  Agrupamento por categoria (map):")
	imprimirGrupos(agruparPorCategoria(escolhidos))
}

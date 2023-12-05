package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Entidade representa uma linha de dados
type Entidade struct {
	Nome      string
	Idade     int
	Pontuacao int
}

// PorNome implementa a interface sort.Interface para ordenação por nome
type PorNome []Entidade

func (p PorNome) Len() int           { return len(p) }
func (p PorNome) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PorNome) Less(i, j int) bool { return p[i].Nome < p[j].Nome }

// PorIdade implementa a interface sort.Interface para ordenação por idade
type PorIdade []Entidade

func (p PorIdade) Len() int           { return len(p) }
func (p PorIdade) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PorIdade) Less(i, j int) bool { return p[i].Idade < p[j].Idade }

// LerArquivo lê o arquivo CSV e retorna uma lista de Entidade
func LerArquivo(nomeArquivo string) ([]Entidade, error) {
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()

	leitorCSV := csv.NewReader(arquivo)
	linhas, err := leitorCSV.ReadAll()
	if err != nil {
		return nil, err
	}

	var entidades []Entidade
	// Ignorar a primeira linha (cabeçalho)
	for _, linha := range linhas[1:] {
		idade, _ := strconv.Atoi(linha[1])
		pontuacao, _ := strconv.Atoi(linha[2])

		entidade := Entidade{
			Nome:      linha[0],
			Idade:     idade,
			Pontuacao: pontuacao,
		}

		entidades = append(entidades, entidade)
	}

	return entidades, nil
}

// EscreverArquivo escreve uma lista de Entidade em um arquivo CSV
func EscreverArquivo(nomeArquivo string, entidades []Entidade) error {
	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	escritorCSV := csv.NewWriter(arquivo)
	defer escritorCSV.Flush()

	// Escreve o cabeçalho
	cabecalho := []string{"Nome", "Idade", "Pontuacao"}
	escritorCSV.Write(cabecalho)

	// Escreve as linhas
	for _, entidade := range entidades {
		linha := []string{entidade.Nome, strconv.Itoa(entidade.Idade), strconv.Itoa(entidade.Pontuacao)}
		escritorCSV.Write(linha)
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <arquivo-origem.csv> <arquivo-destino.csv>")
		os.Exit(1)
	}

	arquivoOrigem := os.Args[1]
	arquivoDestinoNome := "ordenado_por_nome.csv"
	arquivoDestinoIdade := "ordenado_por_idade.csv"

	// Ler o arquivo de entrada
	entidades, err := LerArquivo(arquivoOrigem)
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo de entrada: %v\n", err)
		os.Exit(1)
	}

	// Ordenar por nome e salvar
	sort.Sort(PorNome(entidades))
	err = EscreverArquivo(arquivoDestinoNome, entidades)
	if err != nil {
		fmt.Printf("Erro ao escrever o arquivo ordenado por nome: %v\n", err)
		os.Exit(1)
	}

	// Ordenar por idade e salvar
	sort.Sort(PorIdade(entidades))
	err = EscreverArquivo(arquivoDestinoIdade, entidades)
	if err != nil {
		fmt.Printf("Erro ao escrever o arquivo ordenado por idade: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Processamento concluído com sucesso.")
}

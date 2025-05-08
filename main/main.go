package main

import (
	"Introducao-Backend/conn"
	"fmt"
)

func main() {

	//Aqui estamos utilizando a funçao de conectar ao banco de dados que esta disponivel no pacote conn que criamos
	// e guardando nossa resposta nas variáveis database (que será a partir dela que manipularemos nosso banco) e err (caso exista um erro)
	database, err := conn.ConnectDatabase()

	//Se algum erro for retornado da funçao, a execuçao da funcao será interrompida
	if err != nil {
		panic(err)
	}


	// *INSERÇÃO DE UM VALOR NO BANCO DE DADOS*

	// Através da variavel database nós podemos chamar a função Prepare()
	// A funcao Prepare() permite a criação de uma query SQL e "prepara" ela para ser utilizada
	// No lugar de colocar os valores direto no texto (o que seria inseguro), ela nos permite usar símbolos $1, $2, que serão substituídos posteriormente
	// A função retorna um objeto preparado para execução (query) e um erro (_ que está sendo ignorado)
	query, _ := database.Prepare("INSERT INTO aluno(nome, curso) VALUES ($1, $2)")

	// A partir do objeto query podemos acessar a função Exec
	// Ela irá preencher os símbolos da nossa query preparada anteriormente de acordo com os valores que passamos para ela
	// $1 vira "Denecley" e $2 vira "Engenharia de software"
	// E depois de fazer isso, a função executa a nossa query no banco de dados
	query.Exec("Denecley", "Engenharia de software")


	// *BUSCA DE TODOS AS LINHAS NA TABELA alunos NO BANCO DE DADOS*

	// Novamente a partir de database podemos chamar a função Query()
	// Ela executa a query que foi passada para ela e retorna todos os valores encontrados (rows) e um erro (_ que está sendo ignorado)
	rows, _ := database.Query("SELECT id, nome, curso FROM aluno")

	// Como rows representa todas as linhas do resultado da query, precisamos percorrer cada uma delas para manipular seus valores
	// A função Next() verifica se existe uma próxima linha dentro de rows. Se sim, ele retorna true e o loop continua. Se não, retorna false e para o loop
	// Além disso, ela também prepara a linha verificada para ser utilizada por uma função Scan() no decorrer do código
	for rows.Next() {

		// Variáveis para receber os valores de cada coluna da linha
		var nome, curso string
		var id int

		// A função Scan() copia os valores de cada coluna da linha analisada para as variáveis que passamos a ela
		// A ordem dos parâmetros do Scan() precisa bater com a ordem das colunas do SELECT
		rows.Scan(
			&id,
			&nome,
			&curso,
		)

		// Imprime os dados lidos da linha atual
		fmt.Println("---------------------------------------------------------")
		fmt.Printf("ID: %d\nNome: %s\nCurso: %s\n", id, nome, curso)
	}


	// *ATUALIZAÇÃO DE UM VALOR NO BANCO DE DADOS*

	// Preparamos a query SQL com um SET (atualiza o campo nome onde o id for igual a um valor)
	updateQuery, _ := database.Prepare("UPDATE aluno SET nome = $1 WHERE id = $2")

	// Executamos a query substituindo os símbolos $1 e $2 pelos valores reais
	// Neste exemplo: o aluno com ID 1 terá o nome atualizado para "Denecley Atualizado"
	updateQuery.Exec("Denecley Atualizado", 1)

	fmt.Println("Aluno com ID 1 atualizado com sucesso.")


	// *REMOÇÃO DE UM VALOR DO BANCO DE DADOS*

	// Preparamos a query SQL com DELETE onde o ID for igual ao passado
	deleteQuery, _ := database.Prepare("DELETE FROM aluno WHERE id = $1")

	// Executamos a query passando o ID que será deletado
	// Neste exemplo: o aluno com ID 2 será removido
	deleteQuery.Exec(2)

	fmt.Println("Aluno com ID 2 removido com sucesso.")

}

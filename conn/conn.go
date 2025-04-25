package conn

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

//Variaveis com informacoes necessarias para conexao com o banco de dados
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "introducao"
)

//funcao para conectar com o banco de dados, que não recebe parametros, e retorna um ponteiro do tipo sql.DB e um erro
func ConnectDatabase() (*sql.DB, error) {

	//Formatacao da string que vai ser usada para conectar com o banco
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//A funcao sql.Open() abrirá uma conexao com o banco de dados de acordo com uma string com suas informacoes (no caso a variavel psqlInfo)
	// e com o nome do driver do tipo do banco de dados (no caso o "postgres" que instalamos no projeto e está sendo importado com "github.com/lib/pq")
	db, err := sql.Open("postgres", psqlInfo)

	// Se a conexao com a database der certo, a variável err vai ser igual a "nil", se existir algum erro ela vai ser diferente disso
	// e entrará nesse bloco de código
	if err != nil {

		//A funcao panic() para a execucao da sua funcao. No nosso caso, a execucao da funcao ConnectDatabase() vai parar quando err for 
		// diferente de "nil"
		panic(err)
	}

	// A funcao Ping() verifica se o banco de dados que iniciamos anteriormente (na variavel db) ainda está ativo
	err = db.Ping()

	// Se existir algum erro nessa verificacao, a execucao da funcao será interrompida
	if err != nil {
		panic(err)
	}

	//Caso tudo dê certo, será imprimido no terminal essa mensagem, e retornaremos o banco de dados que criamos e um erro vazio
	fmt.Println("Connected to " + dbname)

	return db, nil
}
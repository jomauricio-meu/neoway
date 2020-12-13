package banco

import (
	"database/sql"
	"fmt"
	"log"
	"neomain/model"
	"neomain/util"
	"runtime"
	"sync"

	_ "github.com/lib/pq"
)

var (
	//fila
	wg sync.WaitGroup
	//limite de conexão.
	concurrencyLevel = runtime.NumCPU() * 8
	//conexão stmt
	StmtMain *sql.Stmt
)

//configurações banco.
const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "NEO"
	dbtable  = "neoway"
)

/*
realiza a conexão do com postgres e
define limites de numero de conexões.
*/
func ConBanco() *sql.DB {
	var db *sql.DB
	var err error

	//string com os dados de conexão
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//efetua a conexão
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalln(err)
	}
	//define o maximo de conexões no banco.
	db.SetMaxIdleConns(concurrencyLevel)

	return db
}

func FechaDb(db *sql.DB) {
	db.Close()
}

func FechaStmt() {
	StmtMain.Close()
}

func Stmt(db *sql.DB, s string) {

	var err error
	StmtMain, err = db.Prepare(s)
	if err != nil {
		log.Fatalln(err)
	}

}

func Persistenca(s []model.Dados) {
	var n int

	n = 0
	for n <= len(s) {
		for k := 0; k < concurrencyLevel; k++ {
			if n <= (len(s) - 1) {
				//adiciona na fila
				wg.Add(1)
				go Grava(s[n])
				n++
			}
		}
		//espera a fila liberar espaço
		wg.Wait()
		if n == len(s) {
			break
		}
	}
}

func Grava(dados model.Dados) {
	defer wg.Done()

	_, err := StmtMain.Exec(dados.Cpf,
		dados.Privado,
		dados.Incompleto,
		util.ValidaNullString(dados.UltimaCompra),
		util.ValidaNullFloat(dados.TicketMedio),
		util.ValidaNullFloat(dados.UltimoTicket),
		dados.LojaFrequente,
		dados.UltimaLoja)

	if err != nil {
		log.Fatalln(err)
	}
}

func GravaDb() []model.Dados {
	var db *sql.DB
	var sDados []model.Dados

	sDados = model.ImportaTxt()

	//conexão banco
	db = ConBanco()

	//seta a query no stmt
	Stmt(db, "INSERT INTO "+dbtable+" (cpf, privado, incompleto, ultimaCompra, ticketMedio, ultimoTicket, lojaFrequente, ultimaLoja) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")

	//persistencia no banco
	Persistenca(sDados)

	//fecha a conexão com o banco
	FechaDb(db)

	//fecha stmt
	FechaStmt()

	return sDados
}

func DeletaTabela() {
	var db *sql.DB

	//conexão banco
	db = ConBanco()

	//roda query
	db.Query("DROP TABLE " + dbtable)

	//fecha a conexão com o banco
	FechaDb(db)

}
func CriaTabela() {
	var db *sql.DB

	//conexão banco
	db = ConBanco()

	//roda query
	db.Query("CREATE TABLE " + dbtable + "(id serial PRIMARY KEY,cpf VARCHAR (50) NOT NULL,privado INTEGER,incompleto INTEGER,ultimaCompra DATE,ticketMedio NUMERIC (10, 2),ultimoTicket NUMERIC (10, 2),lojaFrequente VARCHAR (50),ultimaLoja VARCHAR (50));")

	//fecha a conexão com o banco
	FechaDb(db)

}

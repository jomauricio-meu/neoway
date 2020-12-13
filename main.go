package main

import (
	"encoding/json"
	"log"
	"neomain/banco"
	"neomain/model"
	"net/http"
)

func ImportTxt(w http.ResponseWriter, r *http.Request) {
	var sDados []model.Dados
	//realiza a importação do txt e gravação no banco de dados.
	sDados = banco.GravaDb()
	//retorna o json com todos os registros.
	json.NewEncoder(w).Encode(sDados)
}

func DeletaTabela(w http.ResponseWriter, r *http.Request) {

	//realiza o drop da tabela.
	banco.DeletaTabela()

	json.NewEncoder(w).Encode("Tabela Deletada")
}

func CriaTabela(w http.ResponseWriter, r *http.Request) {

	//realiza o drop da tabela.
	banco.CriaTabela()

	json.NewEncoder(w).Encode("Tabela Criada")
}

func handleRequests() {
	http.HandleFunc("/importa", ImportTxt)
	http.HandleFunc("/deletatabela", DeletaTabela)
	http.HandleFunc("/criatabela", CriaTabela)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}

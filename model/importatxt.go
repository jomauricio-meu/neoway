package model

import (
	"bufio"
	"neomain/util"
	"os"
	"strings"
)

func ImportaTxt() []Dados {

	var sDados []Dados

	f, _ := os.Open("base/base_teste.txt")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)

	if !scanner.Scan() {
		return sDados
	}

	sDados = buffer(*scanner)

	return sDados
}

func buffer(scanner bufio.Scanner) []Dados {

	var dados Dados
	var sDados []Dados

	for scanner.Scan() {

		line := scanner.Text()
		linhaArray := strings.Fields(line)

		if util.IsCPF(linhaArray[0]) && util.IsCNPJ(linhaArray[6]) && util.IsCNPJ(linhaArray[7]) {
			dados.Cpf = util.LimpaEspecial(linhaArray[0])
			dados.Privado = linhaArray[1]
			dados.Incompleto = linhaArray[2]
			dados.UltimaCompra = linhaArray[3]
			dados.TicketMedio = linhaArray[4]
			dados.UltimoTicket = linhaArray[5]
			dados.LojaFrequente = util.LimpaEspecial(linhaArray[6])
			dados.UltimaLoja = util.LimpaEspecial(linhaArray[7])

			sDados = append(sDados, dados)
		}

	}
	return sDados
}

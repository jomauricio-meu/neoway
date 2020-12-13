package model

type Dados struct {
	ID            int    `json:"id,omitempty"`
	Cpf           string `json:"cpf,omitempty"`
	Privado       string `json:"privado,omitempty"`
	Incompleto    string `json:"imcompleto,omitempty"`
	UltimaCompra  string `json:"ultimacompra,omitempty"`
	TicketMedio   string `json:"ticketmedio,omitempty"`
	UltimoTicket  string `json:"ultimoticket,omitempty"`
	LojaFrequente string `json:"lojafrequente,omitempty"`
	UltimaLoja    string `json:"ultimaloja,omitempty"`
}
